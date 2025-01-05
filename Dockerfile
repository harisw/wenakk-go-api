# Start from the official Golang image
FROM golang:1.23-bookworm

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download && go mod verify

RUN go install github.com/air-verse/air@latest

COPY ./ /app/

# Expose the application port
# EXPOSE 8080

# Command to run the application with Air
# CMD ["air -buildvcs=false"]