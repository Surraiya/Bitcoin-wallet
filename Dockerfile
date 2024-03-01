# Use the official golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the source code from the current directory to the working directory inside the container
COPY . .

# Build the Go application
RUN go build -o main .

# Run tests
RUN go test ./...

# Expose the port that your application listens on
EXPOSE 8083

# Run the Go application
CMD ["./main"]
