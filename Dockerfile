# Use an official Golang image as the base image for building the binary
FROM golang:1.20-alpine AS builder

# Install git as it is needed for downloading some dependencies
RUN apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to cache dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o Codewise-CLI ./main.go

# Use a smaller base image for running the binary
FROM alpine:3.18

# Set up working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/Codewise-CLI .

# EXPOSE 8080

# Run the binary
ENTRYPOINT ["./Codewise-CLI"]
