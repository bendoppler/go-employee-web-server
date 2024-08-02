# Stage 1: Build the Go application
FROM golang:1.20-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
