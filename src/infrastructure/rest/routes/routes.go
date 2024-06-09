// Package routes contains all routes of the application
package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	// swaggerFiles for documentation
	_ "github.com/minlebay/pausalac/docs"
	"github.com/minlebay/pausalac/src/infrastructure/rest/adapter"
)

// Security is a struct that contains the security of the application
// @SecurityDefinitions.jwt
type Security struct {
	Authorization string `header:"Authorization" json:"Authorization"`
}

// @title Golang Pausalac
// @version 1.2
// @description Documentation's Golang Pausalac
// @termsOfService http://swagger.io/terms/

// @contact.name Ilshat Minnibaev
// @contact.url https://github.com/minlebay
// @contact.email ilshatminnibaev@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// ApplicationV1Router is a function that contains all routes of the application
// @host localhost:8080
// @BasePath /v1
func ApplicationV1Router(router *gin.Engine, db *mongo.Database) {
	api := router.Group("/api/v1")
	{
		// Documentation Swagger
		{
			api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		}

		AuthRoutes(api, adapter.AuthAdapter(db))
		UserRoutes(api, adapter.UserAdapter(db))
	}
}
