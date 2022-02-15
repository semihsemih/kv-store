FROM golang:1.17-alpine AS builder

WORKDIR /app
RUN apk add gcc g++ ca-certificates --no-cache

COPY go.mod .
RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w -extldflags "-static"' ./cmd/kv-store

FROM alpine:latest

WORKDIR /app

EXPOSE 8080

ENTRYPOINT ["/app/kv-store"]

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder /app/kv-store /app/kv-store