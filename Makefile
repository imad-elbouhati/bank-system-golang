postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin -v /home/vboxuser/postgres-bank-volume/data:/var/lib/postgresql/data -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root bank

dropdb:
	docker exec -it postgres12 dropdb bank

migrateup:	
	migrate -path db/migration -verbose -database "postgres://root:admin@localhost:5432/bank?sslmode=disable" up

migratedown:
	migrate -path db/migration -verbose -database "postgres://root:admin@localhost:5432/bank?sslmode=disable" down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...
	
.PHONY: postgres createdb dropdb migrateup migratedown sqlc