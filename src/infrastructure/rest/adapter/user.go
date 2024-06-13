// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	"go.mongodb.org/mongo-driver/mongo"
	userService "pausalac/src/application/usecases"
	userRepository "pausalac/src/infrastructure/repository"
	userController "pausalac/src/infrastructure/rest/controllers/user"
)

// UserAdapter is a function that returns a user controller
func UserAdapter(db *mongo.Database) *userController.UserController {
	uRepository := userRepository.UserRepository{Collection: db.Collection("users")}
	service := userService.UserService{Repo: &uRepository}
	return &userController.UserController{Service: &service}
}
