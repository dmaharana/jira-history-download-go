# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Binary names
BINARY_NAME=jira-history-download
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_WINDOWS=$(BINARY_NAME).exe

# Build directory
BUILD_DIR=build

# Main package path
MAIN_PACKAGE=./cmd/jirahistorydownload

# Version information
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME=$(shell date +%FT%T%z)
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)"

.PHONY: all build clean test deps build-linux build-windows

all: clean deps test build

build: build-linux build-windows

deps:
	$(GOMOD) download
	$(GOMOD) verify

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

# Build for Linux
build-linux:
	mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_UNIX) $(MAIN_PACKAGE)

# Build for Windows
build-windows:
	mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_WINDOWS) $(MAIN_PACKAGE)

# Run the application
run:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_UNIX) $(MAIN_PACKAGE)
	./$(BUILD_DIR)/$(BINARY_UNIX)

# Cross compilation
.PHONY: build-all
build-all: build-linux build-windows

# Help target
help:
	@echo "Available targets:"
	@echo "  all          - Clean, get dependencies, run tests, and build for all platforms"
	@echo "  build        - Build for all platforms"
	@echo "  clean        - Remove build artifacts"
	@echo "  deps         - Get dependencies"
	@echo "  test         - Run tests"
	@echo "  build-linux  - Build for Linux"
	@echo "  build-windows- Build for Windows"
	@echo "  run          - Build and run the application"
	@echo "  help         - Show this help message"
