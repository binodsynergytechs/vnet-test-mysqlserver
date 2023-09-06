# Build Stage
FROM golang:1.19-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . /app/

# Install gcc and other dependencies
RUN apk add --no-cache gcc musl-dev

# Build the Go app
RUN go build -o main main.go

# Serve Stage
FROM alpine:latest

WORKDIR /app

ENV GOENV=dev

COPY --from=builder /app/main .

COPY --from=builder /app/cmd/ cmd/
COPY --from=builder /app/config/ config/

EXPOSE 8092

CMD ["./main"]