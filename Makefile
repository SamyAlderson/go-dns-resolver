# Build and run the DNS resolver
bin := go-dns-resolver
version := $(shell go version)

# Check if Go is installed
check-go:
	@echo "$(version)" | grep -q "go" || { echo "Error: Go not found"; exit 1; }

# Download dependencies
get-deps:
	@go mod tidy
	@go mod download

# Build the binary
build: check-go get-deps
	@go build -o $(bin) .

# Run the binary
run: build
	@./$(bin)

# Test the binary
test: build
	@go test -v ./...

# Clean up
clean:
	@rm -f $(bin)

# Install the binary
install: build
	@cp $(bin) /usr/local/bin/

# Uninstall the binary
uninstall:
	@rm /usr/local/bin/go-dns-resolver

# Update the module
update:
	@go mod tidy

.PHONY: check-go get-deps build run test clean install uninstall update