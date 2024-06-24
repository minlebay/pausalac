// Package routes contains all routes of the application
package routes

import (
	"github.com/gin-gonic/gin"
	userController "pausalac/src/infrastructure/rest/controllers"
	"pausalac/src/infrastructure/rest/middlewares"
)

func UserRoutes(router *gin.RouterGroup, controller *userController.UserController) {
	routerUser := router.Group("/users")
	routerUser.POST("/createadmin", controller.CreateAdmin)

	routerUser.Use(middlewares.AuthJWTMiddleware())
	{
		routerUser.POST("/", controller.Create)
		routerUser.GET("/:id", controller.GetById)
		routerUser.GET("/", controller.GetAll)
		routerUser.PUT("/:id", controller.Update)
		routerUser.DELETE("/:id", controller.Delete)
	}
}
