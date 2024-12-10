# Dockerfile for building and running the toy-service microservice
# Uses a multi-stage build for smaller images and best practices

FROM golang:1.20-alpine AS builder

WORKDIR /app
COPY . .

# Download dependencies
RUN go mod download

# Build the binary
RUN go build -o toy-service ./cmd/server

# Production image
FROM alpine:3.18
WORKDIR /app

COPY --from=builder /app/toy-service /app/

# Use non-root user for security best practices
RUN adduser -D toyuser
USER toyuser

EXPOSE 8080
ENTRYPOINT ["./toy-service"]