# --- Stage 1: Build ---
FROM golang:1.24.4-alpine AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build your app (main.go is in cmd/)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

# --- Stage 2: Runtime ---
FROM alpine:3.20

WORKDIR /app

# Copy the compiled binary
COPY --from=builder /app/main .

# Copy your SQLite database folder
COPY db ./db

# Expose app port
EXPOSE 8080

CMD ["./main"]
