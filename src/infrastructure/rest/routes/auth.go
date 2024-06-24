package routes

import (
	"github.com/gin-gonic/gin"
	authController "pausalac/src/infrastructure/rest/controllers/auth"
)

func AuthRoutes(router *gin.RouterGroup, controller *authController.Controller) {
	routerAuth := router.Group("/auth")
	{
		routerAuth.POST("/login", controller.Login)
		routerAuth.POST("/access-token", controller.GetAccessTokenByRefreshToken)
	}
}
