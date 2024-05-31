up:
	docker-compose up -d

down:
	docker-compose down

log-api:
	docker-compose logs -f api

up-build:
	docker-compose up -d --build

migrate_up:
	docker-compose exec api migrate -path /go/src/app/internal/db/migrations -database "mysql://user:pass@tcp(db:3306)/local" up

migrate_down:
	docker-compose exec api migrate -path /go/src/app/internal/db/migrations -database "mysql://user:pass@tcp(db:3306)/local" down
