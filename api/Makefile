LOCAL_BIN:=$(CURDIR)/bin

all: test swag buildchat chat buildapi api

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

install:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	
get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u google.golang.org/grpc
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2

generate:
	make generate-service-follow

generate-service-follow:
	mkdir -p followservice/pkg/followservice
	sudo protoc --proto_path followservice/followerservice --proto_path vendor.protogen \
	--go_out=followservice/pkg/followservice --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=followservice/pkg/followservice --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--grpc-gateway_out=followservice/pkg/followservice --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
	followservice/followerservice/follow.proto

run-service-follow:
	@echo "Building follow service..."
	go build -o ./build/service/follow ./followservice/cmd
	@echo "Running follow service..."
	./build/service/follow

vendor-proto:
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi

.PHONY: all test swag buildapi buildchat chat clean run
