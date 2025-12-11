.PHONY: build run test clean help fmt lint

# Build variables
BINARY_NAME=little-share
GO=go

# Default target
help:
	@echo "Little Share - Modern File Sharing Service"
	@echo ""
	@echo "Available targets:"
	@echo "  make build       - Build the application"
	@echo "  make run         - Run the application"
	@echo "  make test        - Run tests"
	@echo "  make fmt         - Format code"
	@echo "  make lint        - Run linter (requires golangci-lint)"
	@echo "  make clean       - Clean build artifacts"
	@echo "  make dev         - Build and run for development"

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@$(GO) build -o $(BINARY_NAME) main.go
	@echo "Build complete: $(BINARY_NAME)"

# Run the application
run: build
	@echo "Running $(BINARY_NAME)..."
	@./$(BINARY_NAME)

# Development: build and run with hot reload capability
dev:
	@echo "Building for development..."
	@$(GO) build -o $(BINARY_NAME) main.go
	@echo "Starting $(BINARY_NAME) in development mode..."
	@./$(BINARY_NAME)

# Run tests
test:
	@echo "Running tests..."
	@$(GO) test -v -race -coverprofile=coverage.out ./...
	@$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Test complete. Coverage report: coverage.html"

# Format code
fmt:
	@echo "Formatting code..."
	@$(GO) fmt ./...
	@echo "Format complete"

# Lint code (requires golangci-lint)
lint:
	@echo "Linting code..."
	@$(GO) vet ./...
	@echo "Lint complete"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -f $(BINARY_NAME)
	@rm -f $(BINARY_NAME).exe
	@rm -f coverage.out coverage.html
	@echo "Clean complete"

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@$(GO) mod download
	@$(GO) mod tidy
	@echo "Dependencies installed"

# View module information
modgraph:
	@$(GO) mod graph

# Check for security issues in dependencies
check-vulns:
	@echo "Checking for vulnerabilities..."
	@$(GO) list -json -m all | nancy sleuth
