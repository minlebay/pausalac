package routes

import (
	"github.com/gin-gonic/gin"
	customerController "pausalac/src/infrastructure/rest/controllers/customer"
	"pausalac/src/infrastructure/rest/middlewares"
)

func CustomerRoutes(router *gin.RouterGroup, controller *customerController.CustomerController) {
	r := router.Group("/customers")
	r.Use(middlewares.AuthJWTMiddleware())

	r.POST("/", controller.Create)
	r.GET("/:id", controller.GetByID)
	r.GET("/", controller.GetAll)
	r.PUT("/:id", controller.Update)
	r.DELETE("/:id", controller.Delete)
}
