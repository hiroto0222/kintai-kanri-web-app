package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hiroto0222/kintai-kanri-web-app/config"
	"github.com/hiroto0222/kintai-kanri-web-app/controllers"
	db "github.com/hiroto0222/kintai-kanri-web-app/db/sqlc"
	"github.com/hiroto0222/kintai-kanri-web-app/routes"
	"github.com/hiroto0222/kintai-kanri-web-app/token"
)

var (
	AuthController     controllers.AuthController
	EmployeeController controllers.EmployeeController
	ClockInController  controllers.ClockInController
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
	EmployeeController = *controllers.NewEmployeeController(config, store, tokenMaker)
	ClockInController = *controllers.NewClockInController(config, store, tokenMaker)

	// setup routers
	server.setupRouter()

	return server, nil
}

// setupRouter sets up the HTTP router for all api endpoints
func (server *Server) setupRouter() {
	router := gin.Default()

	// setup cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{server.Config.Origin},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 15 * time.Second,
	}))

	apiRoutes := router.Group("/api")

	apiRoutes.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Server is running!"})
	})

	AuthRoutes := routes.NewAuthRoutes(AuthController)
	AuthRoutes.AuthRoute(apiRoutes)

	EmployeeRoutes := routes.NewEmployeeRoutes(EmployeeController)
	EmployeeRoutes.EmployeeRoute(apiRoutes)

	ClockInRoutes := routes.NewClockInRoutes(ClockInController)
	ClockInRoutes.ClockInRoute(apiRoutes)

	server.Router = router
}

// start runs the HTTP server on config port
func (server *Server) Start() error {
	return server.Router.Run(":" + server.Config.Port)
}
