// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	"go.mongodb.org/mongo-driver/mongo"
	service "pausalac/src/application/usecases"
	authService "pausalac/src/application/usecases/auth"
	"pausalac/src/domain"
	repo "pausalac/src/infrastructure/repository"
	controller "pausalac/src/infrastructure/rest/controllers/auth"
)

// AuthAdapter is a function that returns a auth controller
func AuthAdapter(db *mongo.Database) *controller.Controller {
	repository := repo.DefaultRepository[domain.User]{Collection: db.Collection("users")}

	entityService := service.EntityService[domain.User]{Repo: &repository}
	userRepository := repo.UserRepository{DefaultRepository: repository}

	userService := service.UserService{
		EntityService: entityService,
		Repo:          userRepository,
	}

	authService := authService.AuthService{UserService: userService}

	return &controller.Controller{AuthService: authService}
}
