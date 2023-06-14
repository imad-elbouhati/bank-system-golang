package main

import (
	"database/sql"
	"log"

	"github.com/imad-elbouhati/bank/api"
	db "github.com/imad-elbouhati/bank/db/sqlc"
	_"github.com/lib/pq"

)



const (
	dbDriver = "postgres"
	dbServer = "postgres://root:admin@localhost:5432/bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)



func main() {

	conn, err := sql.Open(dbDriver, dbServer)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	
	err = server.StartServer(serverAddress)
	if(err != nil) {
		log.Fatal("cannot start server:", err)

	}
}