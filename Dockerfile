# Use the official Golang image as the base image
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app


# Copy the Go modules manifest and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Generate Swagger documentation (ensure `swag` is installed locally in the container)
RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init



# Build the Go application
RUN go build -o main .


FROM alpine:latest


WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/main .


ENV GIN_ENV=production 
ENV	GIN_PORT="8080" 
ENV	GIN_ALLOW_ORIGIN="127.0.0.1"
ENV GIN_MODE=release

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
