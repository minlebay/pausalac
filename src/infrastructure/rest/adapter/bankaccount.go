package adapter

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	service "pausalac/src/application/usecases"
	"pausalac/src/domain"
	repo "pausalac/src/infrastructure/repository"
	controller "pausalac/src/infrastructure/rest/controllers/bankaccount"
)

func BankAccountAdapter(db *mongo.Database) *controller.BankAccountController {
	repo := repo.DefaultRepository[domain.BankAccount]{Collection: db.Collection("bankaccounts")}
	service := service.EntityService[domain.BankAccount, domain.NewBankAccount]{
		Repo: &repo,
		MapperFunc: func(ctx context.Context, newBankAccount *domain.NewBankAccount) *domain.BankAccount {
			return newBankAccount.ToDomainBankAccountMapper()
		},
	}
	return &controller.BankAccountController{Service: &service}
}
