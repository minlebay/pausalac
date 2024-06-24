package routes

import (
	"github.com/gin-gonic/gin"
	bankAccountController "pausalac/src/infrastructure/rest/controllers"
	"pausalac/src/infrastructure/rest/middlewares"
)

func BankAccountRoutes(router *gin.RouterGroup, controller *bankAccountController.BankAccountController) {
	bankAccounts := router.Group("/bankaccounts")
	bankAccounts.Use(middlewares.AuthJWTMiddleware())
	{
		bankAccounts.GET("/", controller.GetAll)
		bankAccounts.GET("/:id", controller.GetById)
		bankAccounts.POST("/", controller.Create)
		bankAccounts.PUT("/:id", controller.Update)
		bankAccounts.DELETE("/:id", controller.Delete)
	}
}
