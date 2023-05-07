package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hiroto0222/kintai-kanri-web-app/controllers"
	"github.com/hiroto0222/kintai-kanri-web-app/middlewares"
)

type EmployeeRoutes struct {
	employeeController controllers.EmployeeController
}

func NewEmployeeRoutes(employeeController controllers.EmployeeController) EmployeeRoutes {
	return EmployeeRoutes{employeeController}
}

func (rc *EmployeeRoutes) EmployeeRoute(rg *gin.RouterGroup) {
	router := rg.Group("/employees").Use(middlewares.AuthMiddleware(rc.employeeController.TokenMaker))
	router.GET("/:id", rc.employeeController.GetEmployee)
}
