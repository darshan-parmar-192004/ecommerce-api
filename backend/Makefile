# Simple Makefile for a Go project

# Build the application
all: build test

build:
	@echo "Building..."
	@go build -o main ./cmd/api

# Run the application
run:
	@go run ./cmd/api

# Create DB container
docker-run:
	@if docker compose up --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up --build; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./database -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main migrate seed

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

# Migration Commands (like boilerplate)

# Run all pending migrations
migrate-up:
	BLUEPRINT_DB_HOST=localhost go run ./cmd/migrate up

# Rollback last migration
migrate-down:
	BLUEPRINT_DB_HOST=localhost go run ./cmd/migrate down

# Re-run last migration
migrate-redo:
	BLUEPRINT_DB_HOST=localhost go run ./cmd/migrate redo

# Show migration status
migrate-status:
	BLUEPRINT_DB_HOST=localhost go run ./cmd/migrate status

# Force database to specific version (for dirty DB recovery)
# Usage: make migrate-force version=2
migrate-force:
	BLUEPRINT_DB_HOST=localhost go run ./cmd/migrate force $(version)

# Drop all migrations (destructive)
migrate-drop:
	BLUEPRINT_DB_HOST=localhost go run ./cmd/migrate drop

# Create new migration files
# Usage: make migrate-create name=add_users_table
migrate-create:
	@migrate create -ext sql -dir internal/migrations -seq $(name)

# Seed database with CSV data (local)
seed:
	BLUEPRINT_DB_HOST=localhost go run ./cmd/seed

# Run seed inside the container
seed-container:
	cd .. && podman-compose exec app ./seed

.PHONY: all build run test clean watch docker-run docker-down itest migrate-up migrate-down migrate-redo migrate-status migrate-force migrate-drop migrate-create seed seed-container
