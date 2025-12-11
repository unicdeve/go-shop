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
	go build -o bin/api cmd/api/main.go

run:
	go run cmd/api/main.go

dev:
	go run cmd/api/main.go

lint:
	golangci-lint run ./...

format:
	@gofmt -s -w .
	@goimports -w .

migrate-up:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/go-shop?sslmode=disable" up

migrate-down:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5432/go-shop?sslmode=disable" down

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