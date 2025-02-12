# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	
	
	@go build -o tmp/main cmd/offgrid/main.go

# Run the application
run:
	@go run cmd/offgrid/main.go

# Database create tables
tables:
	@go run cmd/tables/main.go

# Database create articles
articles:
	@go run cmd/articles/main.go

# Database drop tables
drop:
	@go run cmd/drop/main.go

# Database create superuser
superuser:
	@go run cmd/superuser/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./test... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

.PHONY: all build run tables articles drop superuser test clean watch
