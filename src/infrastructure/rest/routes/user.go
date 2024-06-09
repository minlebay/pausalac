// Package routes contains all routes of the application
package routes

import (
	"github.com/gin-gonic/gin"
	userController "go_gin_api_clean/src/infrastructure/rest/controllers/user"
	"go_gin_api_clean/src/infrastructure/rest/middlewares"
)

// UserRoutes is a function that contains all routes of the user
func UserRoutes(router *gin.RouterGroup, controller *userController.Controller) {
	routerUser := router.Group("/user")
	routerUser.Use(middlewares.AuthJWTMiddleware())
	{
		routerUser.POST("/", controller.NewUser)
		routerUser.GET("/:id", controller.GetUsersByID)
		routerUser.GET("/", controller.GetAllUsers)
		routerUser.PUT("/:id", controller.UpdateUser)
		routerUser.DELETE("/:id", controller.DeleteUser)
	}
}
