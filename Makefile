postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin -d postgres:14.6-alpine
createdb:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres14 dropdb simple_bank
sqlc: 
	sqlc generate

migrateup1:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup2:
	migrate -path db/migration -database "postgresql://postgres:admin@localhost:5432/simplebank?sslmode=disable" -verbose up

test: 
	go test -v -cover ./...
	
.PHONY: postgres createdb dropdb migrateup1 migrateup2