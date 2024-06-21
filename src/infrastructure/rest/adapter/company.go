package adapter

import (
	"go.mongodb.org/mongo-driver/mongo"
	service "pausalac/src/application/usecases"
	"pausalac/src/domain"
	repo "pausalac/src/infrastructure/repository"
	controller "pausalac/src/infrastructure/rest/controllers"
)

func CompanyAdapter(db *mongo.Database) *controller.CompanyController {
	repo := repo.DefaultRepository[domain.Company]{Collection: db.Collection("companies")}
	service := service.EntityService[domain.Company]{Repo: &repo}
	return &controller.CompanyController{Service: &service}
}
