package main

import (
	"database/sql"
	"log"

	"github.com/fauzanfebrian/simplebank/api"
	"github.com/fauzanfebrian/simplebank/config"
	db "github.com/fauzanfebrian/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

func main() {
	config.LoadConfig()

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
