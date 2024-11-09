# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

# Set working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the entire project to the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Create the final image
FROM alpine:latest

# Set environment variables (optional)
# ENV PORT=8080

# Install necessary libraries for running the app (e.g., libc for the Go binary)
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=builder /app/main .

# Copy the .env file (if you have one) into the image
COPY .env .

# Expose the port the app will run on
EXPOSE 5000

# Run the Go application
CMD ["./main"]
