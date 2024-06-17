package routes

import (
	"github.com/gin-gonic/gin"
	serviceController "pausalac/src/infrastructure/rest/controllers/service"
)

// ServiceRoutes defines the routes for the service entity
func ServiceRoutes(router *gin.RouterGroup, controller *serviceController.ServiceController) {
	services := router.Group("/services")
	{
		services.GET("/", controller.GetAll)
		services.GET("/:id", controller.GetByID)
		services.POST("/", controller.Create)
		services.PUT("/:id", controller.Update)
		services.DELETE("/:id", controller.Delete)
	}
}
