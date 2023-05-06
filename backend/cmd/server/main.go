package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/hiroto0222/kintai-kanri-web-app/config"
	"github.com/hiroto0222/kintai-kanri-web-app/controllers"
	db "github.com/hiroto0222/kintai-kanri-web-app/db/sqlc"
	server "github.com/hiroto0222/kintai-kanri-web-app/server"
	_ "github.com/lib/pq" // lib/pqパッケージは直接は使わないが、sql.Open()を呼び出すときに必要
)

var (
	AuthController controllers.AuthController
)

func main() {
	// load config
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config, %v", err)
	}

	// connect to db
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("could not connect to db, %v", err)
	}

	// create db store
	store := db.NewStore(conn)
	fmt.Println("DB connected successfully...")

	// start server
	server := server.NewServer(config, store)
	log.Fatal(server.Start())
}
