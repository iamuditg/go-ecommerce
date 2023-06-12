# Use the official Golang image as the base image
FROM golang:1.17-alpine as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Build the Go binary
RUN go build -o main

# Create a new image based on a minimal Alpine Linux image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage to the current working directory
COPY --from=builder /app/main .

# Copy the necessary files
COPY ./resources/yml/db.yaml /app/products/resources/yml/
COPY schema.sql products/

# Expose the desired port number
EXPOSE 8000

# Run the application
CMD ["./main", "-config=./resources/yml/db.yaml", "-sql=./schema.sql"]


