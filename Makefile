.PHONY: help deps build fmt lint test run clean coverage docker-build docker-run

help:
	@echo "Available targets:"
	@printf "  %-15s %s\n" "deps" "Download Go module dependencies"
	@printf "  %-15s %s\n" "build" "Compile the toy-service binary"
	@printf "  %-15s %s\n" "fmt" "Format all Go source files with gofmt"
	@printf "  %-15s %s\n" "lint" "Run go vet for static analysis"
	@printf "  %-15s %s\n" "test" "Run Go unit and integration tests"
	@printf "  %-15s %s\n" "run" "Execute toy-service locally"
	@printf "  %-15s %s\n" "clean" "Remove build artifacts"
	@printf "  %-15s %s\n" "coverage" "Generate Go coverage profile"
	@printf "  %-15s %s\n" "docker-build" "Build the Docker image"
	@printf "  %-15s %s\n" "docker-run" "Run the Docker container exposing port 8080"

deps:
	@echo "Downloading Go module dependencies..."
	go mod tidy

build: deps
	@echo "Building toy-service..."
	GO111MODULE=on go build -o bin/toy-service ./cmd/server

fmt:
	@echo "Formatting Go source files..."
	gofmt -w $(shell find . -name '*.go' -not -path './vendor/*')

lint: deps
	@echo "Running go vet..."
	GO111MODULE=on go vet ./...

clean:
	@echo "Removing build artifacts..."
	@rm -rf ./bin

test: deps
	@echo "Running tests..."
	GO111MODULE=on go test ./... -v -cover

coverage: deps
	@echo "Generating coverage report..."
	GO111MODULE=on go test ./... -coverprofile=coverage.out

run: build
	@echo "Running toy-service locally..."
	./bin/toy-service

docker-build:
	@echo "Building Docker image..."
	docker build -t toy-service:latest .

docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 --rm toy-service:latest
