package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"pro.qbitty/simplebank/api"
	db "pro.qbitty/simplebank/db/sqlc"
	"pro.qbitty/simplebank/db/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal(err)
	}
}
