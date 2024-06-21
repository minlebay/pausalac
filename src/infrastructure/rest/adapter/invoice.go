package adapter

import (
	"go.mongodb.org/mongo-driver/mongo"
	service "pausalac/src/application/usecases"
	"pausalac/src/domain"
	repo "pausalac/src/infrastructure/repository"
	controller "pausalac/src/infrastructure/rest/controllers"
)

func InvoiceAdapter(db *mongo.Database) *controller.InvoiceController {
	repo := repo.DefaultRepository[domain.Invoice]{Collection: db.Collection("invoices")}
	service := service.EntityService[domain.Invoice]{Repo: &repo}
	return &controller.InvoiceController{Service: &service}
}
