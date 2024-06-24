package routes

import (
	"github.com/gin-gonic/gin"
	serviceController "pausalac/src/infrastructure/rest/controllers"
	"pausalac/src/infrastructure/rest/middlewares"
)

func ServiceRoutes(router *gin.RouterGroup, controller *serviceController.ServiceController) {
	services := router.Group("/services")
	services.Use(middlewares.AuthJWTMiddleware())

	{
		services.GET("/", controller.GetAll)
		services.GET("/:id", controller.GetById)
		services.POST("/", controller.Create)
		services.PUT("/:id", controller.Update)
		services.DELETE("/:id", controller.Delete)
	}
}
