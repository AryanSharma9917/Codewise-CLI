# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the CLI binary
RUN go build -o codewise main.go

# Stage 2: Minimal runtime image
FROM alpine:latest

# Create a non-root user for better security (optional)
RUN adduser -D appuser

# Set working directory
WORKDIR /home/appuser

# Copy the compiled binary from builder stage
COPY --from=builder /app/codewise .

# Set ownership (if using non-root user)
RUN chown -R appuser .

# Use non-root user
USER appuser

# Command to run the CLI
ENTRYPOINT ["./codewise"]
