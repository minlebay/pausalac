package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BankAccount struct {
	Id            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Author        string             `bson:"author" binding:"required" json:"author"`
	AccountNumber string             `bson:"account_number" binding:"required" json:"account_number"`
	BankName      string             `bson:"bank_name" binding:"required" json:"bank_name"`
	SwiftCode     string             `bson:"swift_code" binding:"required" json:"swift_code"`
	IBAN          string             `bson:"iban" binding:"required" json:"iban"`
	Currency      string             `bson:"currency" binding:"required" json:"currency"`
	CreatedAt     primitive.DateTime `bson:"created_at" json:"-"`
	UpdatedAt     primitive.DateTime `bson:"updated_at" json:"-"`
}

type BankAccountService interface {
	GetAll(context.Context) ([]*BankAccount, error)
	GetById(ctx context.Context, id string) (*BankAccount, error)
	Create(ctx context.Context, newBankAccount *BankAccount) (*BankAccount, error)
	Update(ctx context.Context, id string, bankAccount *BankAccount) (*BankAccount, error)
	Delete(ctx context.Context, id string) error
}
