package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroto0222/kintai-kanri-web-app/controllers"
)

type AuthRoutes struct {
	authController controllers.AuthController
}

func NewAuthRoutes(authController controllers.AuthController) AuthRoutes {
	return AuthRoutes{authController}
}

func (rc *AuthRoutes) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")
	router.POST("/register", rc.authController.SignUpEmployee)
	router.POST("/login", rc.authController.SignInEmployee)
}