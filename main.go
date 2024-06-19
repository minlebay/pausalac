package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"pausalac/src/infrastructure/rest/controllers/errors"
	"strings"
	"time"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"pausalac/src/infrastructure/rest/routes"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
// @BasePath /api/v1
func main() {
	router := gin.Default()
	router.Use(limit.MaxAllowed(200))
	router.Use(cors.Default())

	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error in config file: %s", err.Error())
	}

	mongoURI := viper.GetString("Databases.MongoDB.URL")
	databaseName := viper.GetString("Databases.MongoDB.DatabaseLogs")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	db := client.Database(databaseName)

	router.Use(errors.Handler)
	routes.ApplicationV1Router(router, db)

	startServer(router)
}

func startServer(router http.Handler) {
	serverPort := fmt.Sprintf(":%s", viper.GetString("ServerPort"))
	s := &http.Server{
		Addr:           serverPort,
		Handler:        router,
		ReadTimeout:    18000 * time.Second,
		WriteTimeout:   18000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("fatal error description: %s", strings.ToLower(err.Error()))
	}
}
