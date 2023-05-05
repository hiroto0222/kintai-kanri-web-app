package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiroto0222/kintai-kanri-web-app/config"
	db "github.com/hiroto0222/kintai-kanri-web-app/db/sqlc"
	_ "github.com/lib/pq" // lib/pqパッケージは直接は使わないが、sql.Open()を呼び出すときに必要
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
	server := newServer(config, store)
	log.Fatal(server.start())
}

// TODO: Refactor
// 全てのHTTPリクエストを処理するHTTP APIサーバ
type Server struct {
	config config.Config
	store  db.Store    // dbを持つために構造体にする
	router *gin.Engine // 各APIリクエストを正しいハンドラに送信して処理する
}

// newServer creates a new HTTP server
func newServer(config config.Config, store db.Store) *Server {
	server := &Server{
		config: config,
		store:  store,
	}
	server.setupRouter()

	return server
}

// setupRouter sets up the HTTP router for all api endpoints
func (server *Server) setupRouter() {
	router := gin.Default()
	apiRoutes := router.Group("/api")
	apiRoutes.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Server is running!"})
	})

	server.router = router
}

// start runs the HTTP server on config port
func (server *Server) start() error {
	return server.router.Run(":" + server.config.Port)
}
