.PHONY: deps build test run docker-build docker-run

deps:
	@echo "Downloading Go module dependencies..."
	go mod tidy

build: deps
	@echo "Building toy-service..."
	GO111MODULE=on go build -o bin/toy-service ./cmd/server

test: deps
	@echo "Running tests..."
	GO111MODULE=on go test ./... -v -cover

run: build
	@echo "Running toy-service locally..."
	./bin/toy-service

docker-build:
	@echo "Building Docker image..."
	docker build -t toy-service:latest .

docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 --rm toy-service:latest