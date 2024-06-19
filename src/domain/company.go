package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Company struct {
	Id                           primitive.ObjectID `bson:"_id,omitempty"`
	Author                       string             `bson:"author" binding:"required"`
	AgencyId                     string             `bson:"agency_id"`
	Name                         string             `bson:"name" binding:"required"`
	FullName                     string             `bson:"full_name"`
	PIB                          string             `bson:"pib"`
	IdentificationNumber         string             `bson:"identification_number"`
	ForeignExchangeAccountNumber string             `bson:"foreign_exchange_account_number"`
	CallNumber                   string             `bson:"call_number"`
	DateOfRegistration           string             `bson:"date_of_registration"`
	City                         string             `bson:"city"`
	ActivityCodeId               string             `bson:"activity_code_id"`
	MunicipalityId               string             `bson:"municipality_id"`
	Logo                         string             `bson:"logo"`
	StreetAddress                string             `bson:"street_address"`
	StreetNumber                 string             `bson:"street_number"`
	Phone                        string             `bson:"phone"`
	AgencyEmail                  string             `bson:"agency_email"`
	Signature                    string             `bson:"signature"`
	EmploymentType               string             `bson:"employment_type"`
	InvoiceDescription           string             `bson:"invoice_description"`
	CreatedAt                    time.Time          `bson:"created_at"`
	UpdatedAt                    time.Time          `bson:"updated_at"`
	BankAccounts                 []BankAccount      `bson:"bank_accounts"`
}

type NewCompany struct {
	Author                       string
	AgencyId                     string
	Name                         string
	FullName                     string
	PIB                          string
	IdentificationNumber         string
	ForeignExchangeAccountNumber string
	CallNumber                   string
	DateOfRegistration           string
	City                         string
	ActivityCodeId               string
	MunicipalityId               string
	Logo                         string
	StreetAddress                string
	StreetNumber                 string
	Phone                        string
	AgencyEmail                  string
	Signature                    string
	EmploymentType               string
	InvoiceDescription           string
	BankAccounts                 []NewBankAccount
}

func (newCompany *NewCompany) ToDomainCompanyMapper() *Company {

	var bankAccounts []BankAccount
	for _, bankAccount := range newCompany.BankAccounts {
		bankAccounts = append(bankAccounts, *bankAccount.ToDomainBankAccountMapper())
	}

	return &Company{
		Id:                           primitive.NewObjectID(),
		CreatedAt:                    time.Now(),
		UpdatedAt:                    time.Now(),
		Author:                       newCompany.Author,
		AgencyId:                     newCompany.AgencyId,
		Name:                         newCompany.Name,
		FullName:                     newCompany.FullName,
		PIB:                          newCompany.PIB,
		IdentificationNumber:         newCompany.IdentificationNumber,
		ForeignExchangeAccountNumber: newCompany.ForeignExchangeAccountNumber,
		CallNumber:                   newCompany.CallNumber,
		DateOfRegistration:           newCompany.DateOfRegistration,
		City:                         newCompany.City,
		ActivityCodeId:               newCompany.ActivityCodeId,
		MunicipalityId:               newCompany.MunicipalityId,
		Logo:                         newCompany.Logo,
		StreetAddress:                newCompany.StreetAddress,
		StreetNumber:                 newCompany.StreetNumber,
		Phone:                        newCompany.Phone,
		AgencyEmail:                  newCompany.AgencyEmail,
		Signature:                    newCompany.Signature,
		EmploymentType:               newCompany.EmploymentType,
		InvoiceDescription:           newCompany.InvoiceDescription,
		BankAccounts:                 bankAccounts,
	}
}

type CompanyService interface {
	GetAll(context.Context) (*[]Company, error)
	GetByID(ctx context.Context, id string) (*Company, error)
	Create(ctx context.Context, newCompany *NewCompany) (*Company, error)
	Update(ctx context.Context, id string, companyMap map[string]interface{}) (*Company, error)
	Delete(ctx context.Context, id string) error
}
