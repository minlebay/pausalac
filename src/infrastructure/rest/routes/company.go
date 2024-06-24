package routes

import (
	"github.com/gin-gonic/gin"
	companyController "pausalac/src/infrastructure/rest/controllers"
	"pausalac/src/infrastructure/rest/middlewares"
)

func CompanyRoutes(router *gin.RouterGroup, controller *companyController.CompanyController) {
	company := router.Group("/companies")
	company.Use(middlewares.AuthJWTMiddleware())
	{
		company.POST("/", controller.Create)
		company.GET("/:id", controller.GetById)
		company.GET("/", controller.GetAll)
		company.PUT("/:id", controller.Update)
		company.DELETE("/:id", controller.Delete)
	}
}
