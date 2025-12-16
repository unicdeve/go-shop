.PHONY: help build run dev lint format migrate-up migrate-down migrate-create clean test docker-up docker-down

help:
	@echo "Usage: make <target>"
	@echo "Targets:"
	@echo "  build: Build the API"
	@echo "  run: Run the API"
	@echo "  dev: Run the API in development mode"
	@echo "  lint: Run the linter"
	@echo "  format: Format the code"
	@echo "  migrate-up: Run the migrations up"
	@echo "  migrate-down: Run the migrations down"
	@echo "  migrate-create: Create a new migration"
	@echo "  clean: Clean the build artifacts"
	@echo "  docker-up: Start the Docker containers"
	@echo "  docker-down: Stop the Docker containers"

build:
	@echo "Building all binaries...."
	@mkdir -p bin
	@for cmd in cmd/*/; do \
    		if [ -d "$$cmd" ]; then \
    			binary=$$(basename $$cmd); \
    			echo "Building $$binary..."; \
    			go build -o bin/$$binary ./$$cmd; \
    		fi \
    	done

run:
	go run cmd/api/main.go

dev:
	go run cmd/api/main.go

lint:
	golangci-lint run ./...

format:
	@gofmt -s -w .
	@goimports -w .

docs-generate:
	mkdir -p docs
	swag init -g cmd/api/main.go -o docs --parseDependency --parseInternal --exclude .git,docs,docker,db

migrate-up:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5434/go_shop?sslmode=disable" up

migrate-down:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5434/go_shop?sslmode=disable" down

migrate-create:
	migrate create -ext sql -dir db/migrations -seq $(name)

docker-up:
	docker compose -f docker/docker-compose.yml up -d

docker-down:
	docker compose -f docker/docker-compose.yml down

clean:
	rm -f bin/api

test:
	go test ./...

graph-generate:
	@go get github.com/99designs/gqlgen
	@go run github.com/99designs/gqlgen generate