# Use golang:1.17 as the base image
FROM golang:1.20

ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin

# Set the working directory inside the container
WORKDIR /app

# Copy the Go project files into the container
COPY . .

# Run swag init to generate Swagger documentation
RUN go install && go install github.com/swaggo/swag/cmd/swag@latest && swag init

# Build the Go application inside the container
RUN go build -o main .

# Expose the port the application listens on
EXPOSE 8080

# Run the application
CMD ["./main"]