# Use the official Golang 1.19 image as the base image
FROM golang:1.20

WORKDIR /app

# Copy the Go modules and download the dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the project files into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Start the application
CMD ["./main"]