package adapter

import (
	"go.mongodb.org/mongo-driver/mongo"
	customerService "pausalac/src/application/usecases"
	customerRepository "pausalac/src/infrastructure/repository"
	customerController "pausalac/src/infrastructure/rest/controllers/customer"
)

func CustomerAdapter(db *mongo.Database) *customerController.CustomerController {
	cRepository := customerRepository.CustomerRepository{Collection: db.Collection("customers")}
	service := customerService.CustomerService{Repo: &cRepository}
	return &customerController.CustomerController{Service: &service}
}
