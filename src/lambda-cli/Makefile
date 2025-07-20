# Makefile for PoorServerless

.PHONY: build clean test install run-example help

# Build variables
BINARY_NAME=lambda-cli
BUILD_DIR=build
INSTALL_PATH=/usr/local/bin

# Default target
help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build the CLI binary
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) .
	@echo "Binary built: $(BUILD_DIR)/$(BINARY_NAME)"

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@go clean
	@echo "Clean complete"

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

install: build ## Install binary to system PATH
	@echo "Installing $(BINARY_NAME) to $(INSTALL_PATH)..."
	@sudo cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_PATH)/
	@echo "Installation complete"

uninstall: ## Remove binary from system PATH
	@echo "Uninstalling $(BINARY_NAME)..."
	@sudo rm -f $(INSTALL_PATH)/$(BINARY_NAME)
	@echo "Uninstall complete"

run-example: build ## Run example function creation
	@echo "Creating example functions..."
	@mkdir -p examples
	@./$(BUILD_DIR)/$(BINARY_NAME) create-function --name example-python --runtime python --output examples
	@./$(BUILD_DIR)/$(BINARY_NAME) create-function --name example-nodejs --runtime nodejs --output examples
	@./$(BUILD_DIR)/$(BINARY_NAME) create-function --name example-go --runtime go --output examples
	@echo "Example functions created in examples/ directory"

dev: ## Development build and run
	@go run . create-function --name test-func --runtime python --output ./test-output

fmt: ## Format Go code
	@echo "Formatting code..."
	@go fmt ./...

lint: ## Run linter
	@echo "Running linter..."
	@golangci-lint run

mod-check: ## Check for module updates
	@echo "Checking for module updates..."
	@go list -u -m all
