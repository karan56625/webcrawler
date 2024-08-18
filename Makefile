# Variables
APP_NAME := webcrawler
SRC_DIR := ./cmd/$(APP_NAME)
CLIENT_DIR := ./cmd/$(APP_NAME)-client
BUILD_DIR := ./bin
VERSION := 1.0

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

# Build the client and put the binary in the bin directory
build-client:
	@echo "Building $(APP_NAME)-client..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME)-client $(CLIENT_DIR)

run-client: build-client
	@echo "Running $(APP_NAME)-client..."
	@$(BUILD_DIR)/$(APP_NAME)-client

docker-webcrawler:
	docker build -f docker/webcrawler/Dockerfile -t webcrawler:$(VERSION) .

docker-webcrawler-run: docker-webcrawler
	docker run -p 8081:8081 webcrawler:$(VERSION)

docker-webcrawler-push: docker-webcrawler
	docker tag webcrawler:$(VERSION) ghcr.io/karan56625/webcrawler:$(VERSION)
	docker push ghcr.io/karan56625/webcrawler:$(VERSION)

docker-webcrawler-client:
	docker build -f docker/webcrawler-client/Dockerfile -t webcrawler-client:$(VERSION) .

docker-webcrawler-client-push: docker-webcrawler-client
	docker tag webcrawler-client:$(VERSION) ghcr.io/karan56625/webcrawler-client:$(VERSION)
    docker push ghcr.io/karan56625/webcrawler-client:$(VERSION)

#Run all unit tests
test:
	@echo "Running tests..."
	@go test ./...

# Clean up the build directory
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
