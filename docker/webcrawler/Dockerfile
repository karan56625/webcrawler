# Use the official Golang image to build the web crawler
FROM golang:1.23-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Download the Go Modules dependencies
RUN go mod download && go build -o webcrawler ./cmd/webcrawler

# Start a new stage from scratch
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /usr

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/webcrawler .

# Expose port 8080 to the outside world
EXPOSE 8081

# Command to run the executable
CMD ["./webcrawler"]