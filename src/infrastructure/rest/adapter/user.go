// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	"go.mongodb.org/mongo-driver/mongo"
	service "pausalac/src/application/usecases"
	"pausalac/src/domain"
	repo "pausalac/src/infrastructure/repository"
	controller "pausalac/src/infrastructure/rest/controllers"
)

// UserAdapter is a function that returns a user controller
func UserAdapter(db *mongo.Database) *controller.UserController {
	repository := repo.DefaultRepository[domain.User]{Collection: db.Collection("users")}
	service := service.UserService{
		EntityService: service.EntityService[domain.User]{Repo: &repository},
	}
	return &controller.UserController{Service: &service}
}
