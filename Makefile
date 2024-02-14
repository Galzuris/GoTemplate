# Run from .devcontainer

run:
	@go run ./cmd/app/main.go

install:
	@go mod tidy

# Run from Docker environment

up:
	docker compose up -d

down:
	docker compose down --remove-orphans

restart:
	docker compose restart backend
