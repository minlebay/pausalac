package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Company struct {
	Id                           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Author                       string             `bson:"author" binding:"required" json:"author"`
	AgencyId                     string             `bson:"agency_id" json:"agency_id"`
	Name                         string             `bson:"name" binding:"required" json:"name"`
	FullName                     string             `bson:"full_name" json:"full_name"`
	PIB                          string             `bson:"pib" json:"pib"`
	IdentificationNumber         string             `bson:"identification_number" json:"identification_number"`
	ForeignExchangeAccountNumber string             `bson:"foreign_exchange_account_number" json:"foreign_exchange_account_number"`
	CallNumber                   string             `bson:"call_number" json:"call_number"`
	DateOfRegistration           string             `bson:"date_of_registration" json:"date_of_registration"`
	City                         string             `bson:"city" json:"city"`
	ActivityCodeId               string             `bson:"activity_code_id" json:"activity_code_id"`
	MunicipalityId               string             `bson:"municipality_id" json:"municipality_id"`
	Logo                         string             `bson:"logo" json:"logo"`
	StreetAddress                string             `bson:"street_address" json:"street_address"`
	StreetNumber                 string             `bson:"street_number" json:"street_number"`
	Phone                        string             `bson:"phone" json:"phone"`
	AgencyEmail                  string             `bson:"agency_email" json:"agency_email"`
	Signature                    string             `bson:"signature" json:"signature"`
	EmploymentType               string             `bson:"employment_type" json:"employment_type"`
	InvoiceDescription           string             `bson:"invoice_description" json:"invoice_description"`
	CreatedAt                    primitive.DateTime `bson:"created_at" json:"-"`
	UpdatedAt                    primitive.DateTime `bson:"updated_at" json:"-"`
	BankAccounts                 []*BankAccount     `bson:"bank_accounts" json:"bank_accounts"`
}

type CompanyService interface {
	GetAll(context.Context) ([]*Company, error)
	GetById(ctx context.Context, id string) (*Company, error)
	Create(ctx context.Context, company *Company) (*Company, error)
	Update(ctx context.Context, id string, company *Company) (*Company, error)
	Delete(ctx context.Context, id string) error
}
