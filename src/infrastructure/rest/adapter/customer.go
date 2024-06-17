package adapter

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	service "pausalac/src/application/usecases"
	"pausalac/src/domain"
	repo "pausalac/src/infrastructure/repository"
	controller "pausalac/src/infrastructure/rest/controllers/customer"
)

func CustomerAdapter(db *mongo.Database) *controller.CustomerController {
	repo := repo.DefaultRepository[domain.Customer]{Collection: db.Collection("customers")}
	service := service.EntityService[domain.Customer, domain.NewCustomer]{
		Repo: &repo,
		MapperFunc: func(ctx context.Context, newCustomer *domain.NewCustomer) *domain.Customer {
			return newCustomer.ToDomainCustomerMapper()
		},
	}
	return &controller.CustomerController{Service: &service}
}
