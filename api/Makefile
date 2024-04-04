all: clean swag build run

test:
	@echo "Running tests..."
	@go test -v ./... -count=1

swag:
	@echo "Building swagger docs..."
	@swag init -g cmd/api/main.go

buildapi:
	@echo "Building api..."
	@go build -o ./build/api/api ./cmd/api

api: buildapi
	@echo "Running API..."
	@./build/api/api

buildchat:
	@echo "Building chatservice..."
	@go build -o ./build/chats/chatservice ./cmd/chatservice

chat: buildchat
	@echo "Running chatservice..."
	@./build/chats/chatservice

.PHONY: all test swag buildapi buildchat chat clean run 
