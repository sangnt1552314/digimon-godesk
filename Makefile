# Makefile for Digimon GoDesk
.PHONY: help dev prod test build clean logs

# Default target
help:
	@echo "Available commands:"
	@echo "  dev       - Run in development mode"
	@echo "  prod      - Run in production mode"
	@echo "  test      - Run tests"
	@echo "  build     - Build the application"
	@echo "  clean     - Clean up containers and images"
	@echo "  logs      - View application logs"
	@echo "  logs-dev  - View development logs"
	@echo "  logs-prod - View production logs"

# Development environment
dev:
	@echo "Starting development environment..."
	@cp .env.development .env
	docker-compose -f docker-compose.dev.yml up --build

dev-bg:
	@echo "Starting development environment in background..."
	@cp .env.development .env
	docker-compose -f docker-compose.dev.yml up --build -d

# Production environment
prod:
	@echo "Starting production environment..."
	@cp .env.production .env
	docker-compose -f docker-compose.prod.yml up --build

prod-bg:
	@echo "Starting production environment in background..."
	@cp .env.production .env
	docker-compose -f docker-compose.prod.yml up --build -d

# Testing environment
test:
	@echo "Running tests..."
	@cp .env.testing .env
	APP_ENV=testing go test ./...

# Build application locally
build:
	@echo "Building application..."
	go build -o tmp/main ./cmd/main.go

# Run locally without Docker
run-local:
	@echo "Running locally..."
	@cp .env.development .env
	go run ./cmd/main.go

# Clean up
clean:
	@echo "Cleaning up..."
	docker-compose -f docker-compose.dev.yml down --volumes --remove-orphans
	docker-compose -f docker-compose.prod.yml down --volumes --remove-orphans
	docker system prune -f

# View logs
logs:
	@echo "Application logs:"
	@tail -f storage/logs/*.log

logs-dev:
	@echo "Development logs:"
	docker-compose -f docker-compose.dev.yml logs -f

logs-prod:
	@echo "Production logs:"
	docker-compose -f docker-compose.prod.yml logs -f

# Stop services
stop-dev:
	docker-compose -f docker-compose.dev.yml down

stop-prod:
	docker-compose -f docker-compose.prod.yml down

# Restart services
restart-dev:
	docker-compose -f docker-compose.dev.yml restart

restart-prod:
	docker-compose -f docker-compose.prod.yml restart

# View container status
status:
	@echo "Container status:"
	@docker ps -a --filter "name=digimon-godesk"

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy
	go mod download
