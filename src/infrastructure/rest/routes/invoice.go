package routes

import (
	"github.com/gin-gonic/gin"
	invoiceController "pausalac/src/infrastructure/rest/controllers"
	"pausalac/src/infrastructure/rest/middlewares"
)

func InvoiceRoutes(router *gin.RouterGroup, controller *invoiceController.InvoiceController) {
	invoices := router.Group("/invoices")
	invoices.Use(middlewares.AuthJWTMiddleware())
	{
		invoices.GET("/", controller.GetAll)
		invoices.GET("/:id", controller.GetById)
		invoices.POST("/", controller.Create)
		invoices.PUT("/:id", controller.Update)
		invoices.DELETE("/:id", controller.Delete)
	}
}
