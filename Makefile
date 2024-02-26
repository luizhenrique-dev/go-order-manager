include .env
# Variables
APP_NAME=order-manager

# Tasks
default: run

run:
	@echo "Running the application..."
	@go run cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go

build:
	@echo "Building the application..."
	@go build -o $(APP_NAME) cmd/ordersystem/main.go

test:
	@echo "Running tests..."
	@go test ./...

test-coverage:
	@echo "Running tests with coverage..."
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out

.PHONY: run build test test-coverage
