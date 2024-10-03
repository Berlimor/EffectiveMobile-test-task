FROM golang:1.23-alpine as builder

WORKDIR /go/src/app

# Copy go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies; this will be cached as long as go.mod and go.sum don't change
RUN go mod download

# Now copy the rest of the files
COPY . .

# Download install swag
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.3

# Build swagger using swag
RUN swag init -g routers/main.go

# Build the Go application
RUN go build -o main main.go

# Run the Go application
CMD ["./main"]