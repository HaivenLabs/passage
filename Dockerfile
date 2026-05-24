FROM golang:1.22-alpine AS build

WORKDIR /src
COPY backend/go.mod ./backend/go.mod
COPY backend ./backend
WORKDIR /src/backend
RUN go test ./...
RUN go build -o /out/passage-api ./cmd/api

FROM alpine:3.20
RUN addgroup -S passage && adduser -S passage -G passage
USER passage
COPY --from=build /out/passage-api /usr/local/bin/passage-api
EXPOSE 8080
ENTRYPOINT ["passage-api"]
