package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiroto0222/kintai-kanri-web-app/config"
	dbconn "github.com/hiroto0222/kintai-kanri-web-app/db/sqlc"
	_ "github.com/lib/pq" // lib/pqパッケージは直接は使わないが、sql.Open()を呼び出すときに必要
)

var (
	server *gin.Engine
	db     *dbconn.Queries
)

func init() {
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
	db = dbconn.New(conn)

	fmt.Println("DB connected successfully...")

	server = gin.Default()
}

func main() {
	// load config
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config, %v", err)
	}

	// setup router
	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Server is running!"})
	})

	log.Fatal(server.Run(":" + config.Port))
}
