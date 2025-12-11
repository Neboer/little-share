# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o little-share main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Install ca-certificates for HTTPS support
RUN apk --no-cache add ca-certificates

# Copy the binary from builder
COPY --from=builder /app/little-share .

# Copy front-end files
COPY front ./front

# Create volumes
VOLUME ["/app/files"]

# Expose port
EXPOSE 19999

# Set environment variable for Gin
ENV GIN_MODE=release

# Run the application
CMD ["./little-share", "-config", "./little-share-conf.json"]