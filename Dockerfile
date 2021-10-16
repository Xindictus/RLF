FROM golang:alpine

# Set the Current Working Directory inside the container
RUN mkdir -p /api
WORKDIR /api

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# TODO Build multi-stage Dockerfile
# 1. build/compile stage
# 2. run stage
# RUN go build -o ./app ./main.go

# # Install the package
# RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
ENTRYPOINT ["go", "run", "."]