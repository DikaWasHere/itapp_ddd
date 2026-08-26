# --- Stage 1: Build ---
FROM golang:1.26-alpine AS builder
 
WORKDIR /app
 
# Copy go mod files dulu (agar caching lebih efisien)
COPY go.mod go.sum ./
RUN go mod download
 
# Copy semua source code
COPY . .
 
# Build binary
RUN go build -o main .
 
# --- Stage 2: Run ---
FROM alpine:latest
 
WORKDIR /app
 
# Install ca-certificates untuk HTTPS (jika perlu)
RUN apk --no-cache add ca-certificates
 
# Copy binary dari stage sebelumnya
COPY --from=builder /app/main .
 
# Expose port
EXPOSE 8280
 
# Jalankan aplikasi
CMD ["./main"]
