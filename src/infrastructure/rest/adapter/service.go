package adapter

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	service "pausalac/src/application/usecases"
	"pausalac/src/domain"
	repo "pausalac/src/infrastructure/repository"
	controller "pausalac/src/infrastructure/rest/controllers/service"
)

func ServiceAdapter(db *mongo.Database) *controller.ServiceController {
	repo := repo.DefaultRepository[domain.Service]{Collection: db.Collection("services")}
	service := service.EntityService[domain.Service, domain.NewService]{
		Repo: &repo,
		MapperFunc: func(ctx context.Context, newCompany *domain.NewService) *domain.Service {
			return newCompany.ToDomainServiceMapper()
		},
	}
	return &controller.ServiceController{Service: &service}
}
