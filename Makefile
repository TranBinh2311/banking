postgres:
	docker run --name postgres -p 5433:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres simple_bank

migrateup:
	migrate -path db/migration -database "postgres://postgres:postgres@localhost:5433/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://postgres:postgres@localhost:5433/simple_bank?sslmode=disable" -verbose down

dropdb:
	docker exec -it postgres dropdb --username=postgres simple_bank
	
sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb postgresdown sqlc test