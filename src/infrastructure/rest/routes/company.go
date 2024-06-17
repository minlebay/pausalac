package routes

import (
	"github.com/gin-gonic/gin"
	companyController "pausalac/src/infrastructure/rest/controllers/company"
	"pausalac/src/infrastructure/rest/middlewares"
)

// CompanyRoutes defines the routes for the company entity
func CompanyRoutes(router *gin.RouterGroup, controller *companyController.CompanyController) {
	r := router.Group("/companies")
	r.Use(middlewares.AuthJWTMiddleware())

	r.POST("/", controller.Create)
	r.GET("/:id", controller.GetByID)
	r.GET("/", controller.GetAll)
	r.PUT("/:id", controller.Update)
	r.DELETE("/:id", controller.Delete)
}
