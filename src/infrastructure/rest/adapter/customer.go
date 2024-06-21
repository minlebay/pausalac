package adapter

import (
	"go.mongodb.org/mongo-driver/mongo"
	service "pausalac/src/application/usecases"
	"pausalac/src/domain"
	repo "pausalac/src/infrastructure/repository"
	controller "pausalac/src/infrastructure/rest/controllers"
)

func CustomerAdapter(db *mongo.Database) *controller.CustomerController {
	repo := repo.DefaultRepository[domain.Customer]{Collection: db.Collection("customers")}
	service := service.EntityService[domain.Customer]{Repo: &repo}
	return &controller.CustomerController{Service: &service}
}
