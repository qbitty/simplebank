package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"pro.qbitty/simplebank/api"
	db "pro.qbitty/simplebank/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
