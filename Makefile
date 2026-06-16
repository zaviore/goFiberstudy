
.PHONY: dev stage prod build clean help migrate reset

ENV ?= development

help:
	@echo "Available commands:"
	@echo "  make dev         - Run in development mode"
	@echo "  make stage       - Run in staging mode"
	@echo "  make prod        - Run in production mode"
	@echo "  make migrate     - Run database migrations"
	@echo "  make reset       - Reset database (drop all tables and migrate)"
	@echo "  make build       - Build the application"
	@echo "  make clean       - Remove binary files"
	@echo "  make help        - Show this help message"

dev:
	@echo "Running in development mode..."
	@go run main.go --env=development

stage:
	@echo "Running in staging mode..."
	@go run main.go --env=staging

prod:
	@echo "Running in production mode..."
	@go run main.go --env=production

migrate:
	@echo "Running database migrations..."
	@go run main.go --env=$(ENV) --migrate

reset:
	@echo "Resetting database..."
	@read -p "Are you sure? This will delete all data! (yes/no): " confirm; \
	if [ "$$confirm" = "yes" ] || [ "$$confirm" = "y" ]; then \
		go run main.go --env=$(ENV) --reset; \
	else \
		echo "Database reset cancelled."; \
	fi

build:
	@echo "Building application..."
	@go build -o main .

clean:
	@echo "Cleaning up..."
	@rm -f main main.exe
	@echo "Cleanup completed!"
