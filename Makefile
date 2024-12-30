# Define variables for paths
FRONTEND_DIR = frontend
BACKEND_BINARY = app

# Default target, builds everything
all: build

# Install frontend dependencies
install-frontend:
	cd $(FRONTEND_DIR) && npm install

# Build the frontend
build-frontend:
	cd $(FRONTEND_DIR) && npm run build

# Build the Go backend (including the embedded frontend)
build-backend:
	go build -o build/$(BACKEND_BINARY) .

# Clean the dist folder and binary
clean:
	rm -rf $(FRONTEND_DIR)/dist build/$(BACKEND_BINARY)

# Build everything (frontend + backend)
build: install-frontend build-frontend build-backend

# Run the Go application (for development)
run:
	go run main.go

# Phony targets that don't represent actual files
.PHONY: all install-frontend build-frontend build-backend clean build run