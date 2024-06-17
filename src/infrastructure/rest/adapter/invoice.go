package adapter

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	service "pausalac/src/application/usecases"
	"pausalac/src/domain"
	repo "pausalac/src/infrastructure/repository"
	controller "pausalac/src/infrastructure/rest/controllers/invoice"
)

func InvoiceAdapter(db *mongo.Database) *controller.InvoiceController {
	repo := repo.DefaultRepository[domain.Invoice]{Collection: db.Collection("invoices")}
	service := service.EntityService[domain.Invoice, domain.NewInvoice]{
		Repo: &repo,
		MapperFunc: func(ctx context.Context, newInvoice *domain.NewInvoice) *domain.Invoice {
			return newInvoice.ToDomainInvoiceMapper()
		},
	}
	return &controller.InvoiceController{Service: &service}
}
