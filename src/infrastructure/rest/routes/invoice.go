package routes

import (
	"github.com/gin-gonic/gin"
	invoiceController "pausalac/src/infrastructure/rest/controllers/invoice"
	"pausalac/src/infrastructure/rest/middlewares"
)

// InvoiceRoutes defines the routes for the invoice entity
func InvoiceRoutes(router *gin.RouterGroup, controller *invoiceController.InvoiceController) {
	invoices := router.Group("/invoices")
	invoices.Use(middlewares.AuthJWTMiddleware())

	{
		invoices.GET("/", controller.GetAll)
		invoices.GET("/:id", controller.GetByID)
		invoices.POST("/", controller.Create)
		invoices.PUT("/:id", controller.Update)
		invoices.DELETE("/:id", controller.Delete)
	}
}
