package routes

import (
	"github.com/gin-gonic/gin"
	bankAccountController "pausalac/src/infrastructure/rest/controllers/bankaccount"
	"pausalac/src/infrastructure/rest/middlewares"
)

// BankAccountRoutes defines the routes for the bank account entity
func BankAccountRoutes(router *gin.RouterGroup, controller *bankAccountController.BankAccountController) {
	bankAccounts := router.Group("/bankaccounts")
	bankAccounts.Use(middlewares.AuthJWTMiddleware())

	{
		bankAccounts.GET("/", controller.GetAll)
		bankAccounts.GET("/:id", controller.GetByID)
		bankAccounts.POST("/", controller.Create)
		bankAccounts.PUT("/:id", controller.Update)
		bankAccounts.DELETE("/:id", controller.Delete)
	}
}
