postgres:
	docker run --name postgres14 --network bank_network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin -d postgres:14.6-alpine
createdb:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgres14 dropdb simple_bank
sqlc: 
	sqlc generate

newmigration:
	migrate create -ext sql -dir db/migration -seq $(name)

migrateup:
	migrate -path db/migration -database "postgresql://postgres:admin@localhost:5432/simplebank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://postgres:admin@localhost:5432/simplebank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://postgres:admin@localhost:5432/simplebank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://postgres:admin@localhost:5432/simplebank?sslmode=disable" -verbose down 1

test: 
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go  github.com/shariarfaisal/bank/db/sqlc Store

image:
	docker build -t 01822531439/bank:latest .

run:
	docker run --name bank -p 8080:8080 01822531439/bank:latest
	
.PHONY: postgres createdb dropdb migrateup1 migrateup2 sqlc test server mock newmigration migratedown migratedown1