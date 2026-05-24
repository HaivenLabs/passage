package main

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
)

type healthResponse struct {
	Service string `json:"service"`
	Status  string `json:"status"`
}

type server struct {
	databaseURL string
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg := loadConfig()
	srv := &server{databaseURL: cfg.databaseURL}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", srv.health)
	mux.HandleFunc("GET /readyz", srv.ready)

	httpServer := &http.Server{
		Addr:              cfg.httpAddress,
		Handler:           requestLogger(logger, mux),
		ReadHeaderTimeout: 5 * time.Second,
	}

	logger.Info("passage_api_starting", "address", cfg.httpAddress)
	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("passage_api_failed", "error", err)
		os.Exit(1)
	}
}

type config struct {
	httpAddress string
	databaseURL string
}

func loadConfig() config {
	httpAddress := os.Getenv("PASSAGE_HTTP_ADDRESS")
	if httpAddress == "" {
		httpAddress = ":8080"
	}

	databaseURL := os.Getenv("PASSAGE_DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://passage:passage@localhost:5432/passage?sslmode=disable"
	}

	return config{
		httpAddress: httpAddress,
		databaseURL: databaseURL,
	}
}

func (s *server) health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, healthResponse{Service: "passage-api", Status: "ok"})
}

func (s *server) ready(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if err := checkDatabaseEndpoint(ctx, s.databaseURL); err != nil {
		writeJSON(w, http.StatusServiceUnavailable, healthResponse{Service: "passage-api", Status: "database_unavailable"})
		return
	}

	writeJSON(w, http.StatusOK, healthResponse{Service: "passage-api", Status: "ready"})
}

func checkDatabaseEndpoint(ctx context.Context, databaseURL string) error {
	parsedURL, err := url.Parse(databaseURL)
	if err != nil {
		return err
	}

	host := parsedURL.Host
	if host == "" {
		return errors.New("database URL must include host")
	}

	dialer := net.Dialer{}
	conn, err := dialer.DialContext(ctx, "tcp", host)
	if err != nil {
		return err
	}
	return conn.Close()
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func requestLogger(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Info("http_request", "method", r.Method, "path", r.URL.Path, "duration_ms", time.Since(start).Milliseconds())
	})
}
