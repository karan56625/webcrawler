# Variables
APP_NAME := webcrawler
SRC_DIR := ./cmd/$(APP_NAME)
BUILD_DIR := ./bin

# Targets
.PHONY: all build run test clean

# Default target: build, test, and run the application
all: build test run

# Build the application and put the binary in the bin directory
build:
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(SRC_DIR)

# Run the application
run: build
	@echo "Running $(APP_NAME)..."
	@$(BUILD_DIR)/$(APP_NAME)

# Run all unit tests
test:
	@echo "Running tests..."
	@go test ./...

# Clean up the build directory
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
