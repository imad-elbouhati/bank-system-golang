package main

import (
	"database/sql"
	"log"

	"github.com/imad-elbouhati/bank/api"
	db "github.com/imad-elbouhati/bank/db/sqlc"
	"github.com/imad-elbouhati/bank/util"
	_ "github.com/lib/pq"
)






func main() {

	conf,err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot read conf:", err)
	}


	conn, err := sql.Open(conf.DBDriver, conf.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	
	err = server.StartServer(conf.ServerAddress)
	if(err != nil) {
		log.Fatal("cannot start server:", err)

	}
}