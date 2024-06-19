package routes

import (
	"github.com/gin-gonic/gin"
	companyController "pausalac/src/infrastructure/rest/controllers/company"
	"pausalac/src/infrastructure/rest/middlewares"
)

// CompanyRoutes defines the routes for the company entity
func CompanyRoutes(router *gin.RouterGroup, controller *companyController.CompanyController) {
	company := router.Group("/companies")
	company.Use(middlewares.AuthJWTMiddleware())

	company.POST("/", controller.Create)
	company.GET("/:id", controller.GetByID)
	company.GET("/", controller.GetAll)
	company.PUT("/:id", controller.Update)
	company.DELETE("/:id", controller.Delete)
}
