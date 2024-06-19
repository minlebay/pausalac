package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BankAccount struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	Author        string             `bson:"author" binding:"required"`
	AccountNumber string             `bson:"account_number" binding:"required"`
	BankName      string             `bson:"bank_name" binding:"required"`
	SwiftCode     string             `bson:"swift_code" binding:"required"`
	IBAN          string             `bson:"iban" binding:"required"`
	Currency      string             `bson:"currency" binding:"required"`
	CreatedAt     time.Time          `bson:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at"`
}

type NewBankAccount struct {
	Author        string
	AccountNumber string
	BankName      string
	SwiftCode     string
	IBAN          string
	Currency      string
}

func (newBankAccount *NewBankAccount) ToDomainBankAccountMapper() *BankAccount {
	return &BankAccount{
		Id:            primitive.NewObjectID(),
		Author:        newBankAccount.Author,
		AccountNumber: newBankAccount.AccountNumber,
		BankName:      newBankAccount.BankName,
		SwiftCode:     newBankAccount.SwiftCode,
		IBAN:          newBankAccount.IBAN,
		Currency:      newBankAccount.Currency,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}

type BankAccountService interface {
	GetAll(context.Context) (*[]BankAccount, error)
	GetByID(ctx context.Context, id string) (*BankAccount, error)
	Create(ctx context.Context, newBankAccount *NewBankAccount) (*BankAccount, error)
	Update(ctx context.Context, id string, bankAccountMap map[string]interface{}) (*BankAccount, error)
	Delete(ctx context.Context, id string) error
}
