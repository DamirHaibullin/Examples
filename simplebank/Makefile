postgres:
	docker run --name posgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.0-alpine

createdb:
	docker exec -it posgres15 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it posgres15 dropdb

migrationup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose up

migrationdown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose down

sqlc:
	docker run --rm -v D:/Projects/Examples/simplebank:/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/dhaibullin/simplebank/db/sqlc Store

.PHONY: posgres createdb dropdb migrationup migrationdown sqlc test mock
