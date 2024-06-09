// Package routes contains all routes of the application
package routes

import (
	"github.com/gin-gonic/gin"
	authController "go_gin_api_clean/src/infrastructure/rest/controllers/auth"
)

// AuthRoutes is a function that contains all routes of the auth
func AuthRoutes(router *gin.RouterGroup, controller *authController.Controller) {

	routerAuth := router.Group("/auth")
	{
		routerAuth.POST("/login", controller.Login)
		routerAuth.POST("/access-token", controller.GetAccessTokenByRefreshToken)
	}

}