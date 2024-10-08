# Variables
BINARY_NAME=go-noc-system
CMD_DIR=./main.go

# Default target when running `make`
.PHONY: run
run:
	@echo "Running application..."
	go run $(CMD_DIR)

# Build the application
.PHONY: build
build:
	@echo "Building application..."
	go build -o $(BINARY_NAME) $(CMD_DIR)

# Clean the build
.PHONY: clean
clean:
	@echo "Cleaning build..."
	rm -f $(BINARY_NAME)

# Test the application
.PHONY: test
test:
	@echo "Running tests..."
	go test ./...
