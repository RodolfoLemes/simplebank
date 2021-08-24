docker:
	docker compose up

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/master-class?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/master-class?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: run migrateup migratedown sqlc test