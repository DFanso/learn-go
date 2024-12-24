# Use the official Golang image as the base image
FROM golang:1.23

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

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
