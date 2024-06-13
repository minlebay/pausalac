// Package routes contains all routes of the application
package routes

import (
	"github.com/gin-gonic/gin"
	userController "pausalac/src/infrastructure/rest/controllers/user"
	"pausalac/src/infrastructure/rest/middlewares"
)

// UserRoutes is a function that contains all routes of the user
func UserRoutes(router *gin.RouterGroup, controller *userController.UserController) {
	routerUser := router.Group("/users")
	routerUser.POST("/createadmin", controller.CreateAdmin)

	routerUser.Use(middlewares.AuthJWTMiddleware())
	{
		routerUser.POST("/", controller.Create)
		routerUser.GET("/:id", controller.GetByID)
		routerUser.GET("/", controller.GetAll)
		routerUser.PUT("/:id", controller.Update)
		routerUser.DELETE("/:id", controller.Delete)
	}
}
