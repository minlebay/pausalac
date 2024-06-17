package adapter

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	service "pausalac/src/application/usecases"
	"pausalac/src/domain"
	repo "pausalac/src/infrastructure/repository"
	controller "pausalac/src/infrastructure/rest/controllers/company"
)

func CompanyAdapter(db *mongo.Database) *controller.CompanyController {
	repo := repo.DefaultRepository[domain.Company]{Collection: db.Collection("companies")}
	service := service.EntityService[domain.Company, domain.NewCompany]{
		Repo: &repo,
		MapperFunc: func(ctx context.Context, newCompany *domain.NewCompany) *domain.Company {
			return newCompany.ToDomainCompanyMapper()
		},
	}
	return &controller.CompanyController{Service: &service}
}
