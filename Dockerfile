# --- Stage 1: Build tahap pertama ---
FROM golang:1.21-alpine AS builder

# Set working directory di dalam container
WORKDIR /app

# Copy go.mod & go.sum (Karena projectmu di folder test-api, sesuaikan path jika di-build dari root)
# Tetapi asumsinya kita pindahkan struktur 'advanced' menjadi independent.
COPY go.mod go.sum ./
RUN go mod download

# Copy source code ke container
COPY . .

# Build binary di CGO_ENABLED=0 agar fully static (kecuali jika tetap pakai SQLite yang butuh CGO)
# Jika kamu mengubah config DB ke PostgreSQL nanti, gunakan CGO_ENABLED=0
RUN go build -o advanced_api main.go


# --- Stage 2: Minimalis Container ---
FROM alpine:latest

WORKDIR /root/

# Pindahkan hasil build aplikasi dari tahap Builder
COPY --from=builder /app/advanced_api .
COPY --from=builder /app/.env .

# Buka Port
EXPOSE 8082

# Jalankan perintah
CMD ["./advanced_api"]
