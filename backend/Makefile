.PHONY: run dev build test clean fmt lint help db-init

# Variables
BINARY_NAME=learng-api
MAIN_PATH=./cmd/api/main.go

# Default target
help:
	@echo "Available commands:"
	@echo "  make run       - Run the application"
	@echo "  make dev       - Run with hot reload (requires air)"
	@echo "  make build     - Build the binary"
	@echo "  make test      - Run tests"
	@echo "  make clean     - Clean build artifacts"
	@echo "  make fmt       - Format code"
	@echo "  make lint      - Run linter (requires golangci-lint)"
	@echo "  make db-init   - Initialize database"

# Run the application
run:
	@echo "Starting server..."
	@go run $(MAIN_PATH)

# Run with hot reload (requires air: go install github.com/cosmtrek/air@latest)
dev:
	@which air > /dev/null || (echo "Installing air..." && go install github.com/cosmtrek/air@latest)
	@air

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@CGO_ENABLED=1 go build -o $(BINARY_NAME) $(MAIN_PATH)
	@echo "Build complete: $(BINARY_NAME)"

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -cover -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@rm -f coverage.out coverage.html
	@rm -f learng.db learng.db-shm learng.db-wal
	@echo "Clean complete"

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Format complete"

# Run linter (requires golangci-lint)
lint:
	@which golangci-lint > /dev/null || (echo "golangci-lint not found. Install from https://golangci-lint.run/usage/install/" && exit 1)
	@echo "Running linter..."
	@golangci-lint run

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

# Initialize database (creates a fresh database)
db-init:
	@echo "Initializing database..."
	@rm -f learng.db learng.db-shm learng.db-wal
	@go run $(MAIN_PATH) &
	@sleep 2
	@pkill -f "$(BINARY_NAME)" || pkill -f "go run"
	@echo "Database initialized"

# Install development tools
install-tools:
	@echo "Installing development tools..."
	@go install github.com/cosmtrek/air@latest
	@echo "Tools installed. Run 'make dev' for hot reload development"
