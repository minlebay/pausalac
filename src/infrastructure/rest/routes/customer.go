package routes

import (
	"github.com/gin-gonic/gin"
	customerController "pausalac/src/infrastructure/rest/controllers/customer"
	"pausalac/src/infrastructure/rest/middlewares"
)

func CustomerRoutes(router *gin.RouterGroup, controller *customerController.CustomerController) {
	customers := router.Group("/customers")
	customers.Use(middlewares.AuthJWTMiddleware())

	customers.POST("/", controller.Create)
	customers.GET("/:id", controller.GetByID)
	customers.GET("/", controller.GetAll)
	customers.PUT("/:id", controller.Update)
	customers.DELETE("/:id", controller.Delete)
}
