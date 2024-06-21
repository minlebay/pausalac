package adapter

import (
	"go.mongodb.org/mongo-driver/mongo"
	service "pausalac/src/application/usecases"
	"pausalac/src/domain"
	repo "pausalac/src/infrastructure/repository"
	controller "pausalac/src/infrastructure/rest/controllers"
)

func BankAccountAdapter(db *mongo.Database) *controller.BankAccountController {
	repo := repo.DefaultRepository[domain.BankAccount]{Collection: db.Collection("bankaccounts")}
	service := service.EntityService[domain.BankAccount]{Repo: &repo}
	return &controller.BankAccountController{Service: &service}
}
