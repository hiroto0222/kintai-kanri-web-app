package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroto0222/kintai-kanri-web-app/controllers"
	"github.com/hiroto0222/kintai-kanri-web-app/middlewares"
)

type AuthRoutes struct {
	authController controllers.AuthController
}

func NewAuthRoutes(authController controllers.AuthController) AuthRoutes {
	return AuthRoutes{authController}
}

func (rc *AuthRoutes) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.POST(
		"/register",
		middlewares.AuthMiddleware(rc.authController.TokenMaker),
		rc.authController.RegisterEmployee,
	)

	router.POST("/refresh", rc.authController.RefreshAccessToken)
	router.POST("/login", rc.authController.LogInEmployee)
}
