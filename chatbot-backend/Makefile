# Simple Makefile for a Go project

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

run-local:
	@echo "Starting Python backend..."
	@export PY_BE_HOST='127.0.0.1' && sh './scripts/start_py_server.sh' &
	@echo "Starting Go backend..."
	@go run cmd/api/main.go


run-ci: 
	@go run cmd/api/main.go

test:
	@echo "Testing..."
	@go test ./tests -v

clean:
	@echo "Cleaning..."
	@rm -f main

watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean
