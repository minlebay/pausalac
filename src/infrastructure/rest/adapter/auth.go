// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	"go.mongodb.org/mongo-driver/mongo"
	authService "pausalac/src/application/usecases/auth"
	userRepository "pausalac/src/infrastructure/repository"
	authController "pausalac/src/infrastructure/rest/controllers/auth"
)

// AuthAdapter is a function that returns a auth controller
func AuthAdapter(db *mongo.Database) *authController.Controller {
	uRepository := userRepository.UserRepository{Collection: db.Collection("users")}
	service := authService.AuthService{UserRepository: uRepository}
	return &authController.Controller{AuthService: service}
}
