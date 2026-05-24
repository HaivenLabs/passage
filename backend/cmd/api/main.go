package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type healthResponse struct {
	Service string `json:"service"`
	Status  string `json:"status"`
}

type server struct {
	db *sql.DB
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg := loadConfig()
	db, err := openDatabase(cfg.databaseURL)
	if err != nil {
		logger.Error("database_open_failed", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	srv := &server{db: db}

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

func openDatabase(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	return db, nil
}

func (s *server) health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, healthResponse{Service: "passage-api", Status: "ok"})
}

func (s *server) ready(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if err := s.db.PingContext(ctx); err != nil {
		writeJSON(w, http.StatusServiceUnavailable, healthResponse{Service: "passage-api", Status: "database_unavailable"})
		return
	}

	writeJSON(w, http.StatusOK, healthResponse{Service: "passage-api", Status: "ready"})
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
