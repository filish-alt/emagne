.PHONY: run migrate build clean

# Default target
all: migrate run

# Run migrations
migrate:
	@echo "Running database migrations..."
	@go run cmd/migrate/main.go

# Run the application
run:
	@echo "Starting the application..."
	@go run cmd/main.go

# Build the application
build:
	@echo "Building the application..."
	@go build -o bin/emagne cmd/main.go

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@go mod download

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Lint code
lint:
	@echo "Linting code..."
	@golangci-lint run

