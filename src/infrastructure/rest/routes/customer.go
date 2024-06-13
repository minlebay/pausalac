package routes

import (
	"github.com/gin-gonic/gin"
	customerController "pausalac/src/infrastructure/rest/controllers/customer"
	"pausalac/src/infrastructure/rest/middlewares"
)

func CustomerRoutes(router *gin.RouterGroup, controller *customerController.CustomerController) {
	routerCustomer := router.Group("/customers")
	routerCustomer.Use(middlewares.AuthJWTMiddleware())

	routerCustomer.POST("/", controller.Create)
	routerCustomer.GET("/:id", controller.GetByID)
	routerCustomer.GET("/", controller.GetAll)
	routerCustomer.PUT("/:id", controller.Update)
	routerCustomer.DELETE("/:id", controller.Delete)
}
