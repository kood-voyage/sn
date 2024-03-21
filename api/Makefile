all: clean test swag build run

test:
	@echo "Running tests..."
	@go test -v ./... -count=1

swag:
	@echo "Building swagger docs..."
	@swag init -g cmd/api/main.go

build:
	@echo "Building binary..."
	@go build -o ./build/api/api ./cmd/api

clean:
	@echo "Cleaning up..."
	@rm -rf ./build

run: build
	@echo "Running API..."
	@./build/api/api

.PHONY: all test swag build clean run
