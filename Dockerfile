# Build stage
FROM golang:1.23.1 AS builder

WORKDIR /app

# Cache go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o urlshortener ./cmd/urlshortener

# Final stage
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/urlshortener .
EXPOSE 8080

# Run the app
ENTRYPOINT ["./urlshortener"]
