DB_URL=postgres://root:admin@localhost:5432/bank?sslmode=disable

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin -v /home/vboxuser/postgres-bank-volume/data:/var/lib/postgresql/data -d --rm postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root bank

dropdb:
	docker exec -it postgres12 dropdb bank

migrateup:	
	migrate -path db/migration -verbose -database "$(DB_URL)" up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -verbose -database "$(DB_URL)" down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1
	
sqlc:
	sqlc generate

test:
	go test -v -cover ./...
	
server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server