package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	// swaggerFiles for documentation
	_ "pausalac/docs"
	"pausalac/src/infrastructure/rest/adapter"
)

type Security struct {
	Authorization string `header:"Authorization" json:"Authorization"`
}

func ApplicationV1Router(router *gin.Engine, db *mongo.Database) {
	api := router.Group("/api/v1")
	{
		// Documentation Swagger
		{
			api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		AuthRoutes(api, adapter.AuthAdapter(db))
		UserRoutes(api, adapter.UserAdapter(db))
		CustomerRoutes(api, adapter.CustomerAdapter(db))
		CompanyRoutes(api, adapter.CompanyAdapter(db))
		BankAccountRoutes(api, adapter.BankAccountAdapter(db))
		ServiceRoutes(api, adapter.ServiceAdapter(db))
		InvoiceRoutes(api, adapter.InvoiceAdapter(db))
	}
}
