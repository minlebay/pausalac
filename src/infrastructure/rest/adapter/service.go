package adapter

import (
	"go.mongodb.org/mongo-driver/mongo"
	service "pausalac/src/application/usecases"
	"pausalac/src/domain"
	repo "pausalac/src/infrastructure/repository"
	controller "pausalac/src/infrastructure/rest/controllers"
)

func ServiceAdapter(db *mongo.Database) *controller.ServiceController {
	repo := repo.DefaultRepository[domain.Service]{Collection: db.Collection("services")}
	service := service.EntityService[domain.Service]{Repo: &repo}
	return &controller.ServiceController{Service: &service}
}
