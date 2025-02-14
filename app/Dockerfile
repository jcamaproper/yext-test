# Use the official Go image as a builder for a lightweight and secure base
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker caching for dependencies
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application as a statically compiled binary
RUN go build -o sort-service

# Start a new minimal image to reduce the final image size
FROM alpine:latest

# Set the working directory inside the final image
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/sort-service .

# Expose the port your application will run on
EXPOSE 8080

# Command to run the application
CMD ["./sort-service"]