// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	authService "github.com/minlebay/pausalac/src/application/usecases/auth"
	userRepository "github.com/minlebay/pausalac/src/infrastructure/repository/user"
	authController "github.com/minlebay/pausalac/src/infrastructure/rest/controllers/auth"
	"go.mongodb.org/mongo-driver/mongo"
)

// AuthAdapter is a function that returns a auth controller
func AuthAdapter(db *mongo.Database) *authController.Controller {
	uRepository := userRepository.Repository{Collection: db.Collection("users")}
	service := authService.Service{UserRepository: uRepository}
	return &authController.Controller{AuthService: service}
}
