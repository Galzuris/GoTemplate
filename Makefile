# Run from Docker environment

up:
	docker compose up -d

down:
	docker compose down --remove-orphans

restart:
	docker compose restart backend

database:
	docker compose exec database psql -U postgres -d app

migrate-up:
	docker compose exec backend go run -tags migrate ./cmd/app
