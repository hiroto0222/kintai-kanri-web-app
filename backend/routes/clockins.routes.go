package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroto0222/kintai-kanri-web-app/controllers"
	"github.com/hiroto0222/kintai-kanri-web-app/middlewares"
)

type ClockInRoutes struct {
	clockInController controllers.ClockInController
}

func NewClockInRoutes(clockInController controllers.ClockInController) ClockInRoutes {
	return ClockInRoutes{clockInController}
}

func (c *ClockInRoutes) ClockInRoute(rg *gin.RouterGroup) {
	router := rg.Group("/clockins").Use(middlewares.AuthMiddleware(c.clockInController.TokenMaker))

	router.POST("", c.clockInController.CreateClockIn)
	router.GET("/:employee_id", c.clockInController.ListClockInsAndClockOuts)
}
