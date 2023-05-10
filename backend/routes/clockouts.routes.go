package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroto0222/kintai-kanri-web-app/controllers"
	"github.com/hiroto0222/kintai-kanri-web-app/middlewares"
)

type ClockOutRoutes struct {
	clockOutController controllers.ClockOutController
}

func NewClockOutRoutes(clockOutController controllers.ClockOutController) ClockOutRoutes {
	return ClockOutRoutes{clockOutController}
}

func (c *ClockOutRoutes) ClockOutRoute(rg *gin.RouterGroup) {
	router := rg.Group("/clockouts").Use(middlewares.AuthMiddleware(c.clockOutController.TokenMaker))

	router.POST("", c.clockOutController.CreateClockOut)
}
