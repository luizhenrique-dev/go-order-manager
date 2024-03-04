include .env
# Variables
APP_NAME=order-manager

# Tasks
default: run

create_migration:
	@echo "Creating new migration..."
	migrate create -ext=sql -dir=internal/infra/database/migrations -seq order_manager

run:
	@echo "Running the application..."
	@go run cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go

build:
	@echo "Building the application..."
	@go build -o $(APP_NAME) cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go

test:
	@echo "Running tests..."
	@go test ./...

test-coverage:
	@echo "Running tests with coverage..."
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out

docker-build:
	@echo "Building docker image..."
	@docker build -t luizhenrique-dev/go-order-manager:latest .

docker-run:
	@echo "Running docker image..."
	@docker run -p 8000:8000 luizhenrique-dev/go-order-manager:latest

.PHONY: run build test test-coverage create_migration docker-build docker-run
