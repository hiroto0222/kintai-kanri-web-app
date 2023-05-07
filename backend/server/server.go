package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiroto0222/kintai-kanri-web-app/config"
	"github.com/hiroto0222/kintai-kanri-web-app/controllers"
	db "github.com/hiroto0222/kintai-kanri-web-app/db/sqlc"
	"github.com/hiroto0222/kintai-kanri-web-app/routes"
	"github.com/hiroto0222/kintai-kanri-web-app/token"
)

var (
	AuthController controllers.AuthController
)

// 全てのHTTPリクエストを処理するHTTP APIサーバ
type Server struct {
	Config     config.Config
	Store      db.Store    // dbを持つために構造体にする
	Router     *gin.Engine // 各APIリクエストを正しいハンドラに送信して処理する
	TokenMaker token.Maker // トークンを生成する構造体
}

// NewServer creates a new HTTP server
func NewServer(config config.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker, %q", err)
	}

	server := &Server{
		Config:     config,
		Store:      store,
		TokenMaker: tokenMaker,
	}

	// create controllers
	AuthController = *controllers.NewAuthController(config, store, tokenMaker)

	// setup routers
	server.setupRouter()

	return server, nil
}

// setupRouter sets up the HTTP router for all api endpoints
func (server *Server) setupRouter() {
	router := gin.Default()

	apiRoutes := router.Group("/api")

	apiRoutes.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Server is running!"})
	})

	AuthRoutes := routes.NewAuthRoutes(AuthController)
	AuthRoutes.AuthRoute(apiRoutes)

	server.Router = router
}

// start runs the HTTP server on config port
func (server *Server) Start() error {
	return server.Router.Run(":" + server.Config.Port)
}
