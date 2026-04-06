# --- Stage 1: Build ---
FROM golang:1.26-alpine AS builder

WORKDIR /build

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy entire project
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o advanced_api ./advanced/main.go


# --- Stage 2: Runtime ---
FROM alpine:latest

WORKDIR /app

# Install ca-certificates for HTTPS
RUN apk add --no-cache ca-certificates

# Copy binary from builder
COPY --from=builder /build/advanced_api .

EXPOSE 8082

CMD ["./advanced_api"]
