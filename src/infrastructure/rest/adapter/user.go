// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	userService "github.com/minlebay/pausalac/src/application/usecases/user"
	userRepository "github.com/minlebay/pausalac/src/infrastructure/repository/user"
	userController "github.com/minlebay/pausalac/src/infrastructure/rest/controllers/user"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserAdapter is a function that returns a user controller
func UserAdapter(db *mongo.Database) *userController.UserController {
	uRepository := userRepository.Repository{Collection: db.Collection("users")}
	service := userService.Service{Repo: &uRepository}
	return &userController.UserController{Service: &service}
}
