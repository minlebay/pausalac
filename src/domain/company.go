package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Company struct {
	ID                           primitive.ObjectID `bson:"_id,omitempty"`
	UserID                       string             `bson:"user_id" binding:"required"`
	AgencyID                     string             `bson:"agency_id"`
	Name                         string             `bson:"name" binding:"required"`
	FullName                     string             `bson:"full_name"`
	PIB                          string             `bson:"pib"`
	IdentificationNumber         string             `bson:"identification_number"`
	FirstAccountNumber           string             `bson:"first_account_number"`
	SecondAccountNumber          string             `bson:"second_account_number"`
	ForeignExchangeAccountNumber string             `bson:"foreign_exchange_account_number"`
	CallNumber                   string             `bson:"call_number"`
	DateOfRegistration           string             `bson:"date_of_registration"`
	City                         string             `bson:"city"`
	ActivityCodeID               string             `bson:"activity_code_id"`
	MunicipalityID               string             `bson:"municipality_id"`
	EmployedByOtherFirm          string             `bson:"employed_by_other_firm"`
	EmploymentChanged            string             `bson:"employment_changed"`
	Logo                         string             `bson:"logo"`
	StreetAddress                string             `bson:"street_address"`
	StreetNumber                 string             `bson:"street_number"`
	Phone                        string             `bson:"phone"`
	AgencyEmail                  string             `bson:"agency_email"`
	SWIFT                        string             `bson:"swift"`
	IBAN                         string             `bson:"iban"`
	Signature                    string             `bson:"signature"`
	EmploymentType               string             `bson:"employment_type"`
	InvoiceDescription           string             `bson:"invoice_description"`
	CreatedAt                    time.Time          `bson:"created_at"`
	UpdatedAt                    time.Time          `bson:"updated_at"`
	User                         User               `bson:"user"`
	BankAccounts                 []BankAccount      `bson:"bank_accounts"`
}

type NewCompany struct {
	UserID                       string
	AgencyID                     string
	Name                         string
	FullName                     string
	PIB                          string
	IdentificationNumber         string
	FirstAccountNumber           string
	SecondAccountNumber          string
	ForeignExchangeAccountNumber string
	CallNumber                   string
	DateOfRegistration           string
	City                         string
	ActivityCodeID               string
	MunicipalityID               string
	EmployedByOtherFirm          string
	EmploymentChanged            string
	Logo                         string
	StreetAddress                string
	StreetNumber                 string
	Phone                        string
	AgencyEmail                  string
	SWIFT                        string
	IBAN                         string
	Signature                    string
	EmploymentType               string
	InvoiceDescription           string
}

func (newCompany *NewCompany) ToDomainCompanyMapper() *Company {
	return &Company{
		ID:                           primitive.NewObjectID(),
		CreatedAt:                    time.Now(),
		UpdatedAt:                    time.Now(),
		UserID:                       newCompany.UserID,
		AgencyID:                     newCompany.AgencyID,
		Name:                         newCompany.Name,
		FullName:                     newCompany.FullName,
		PIB:                          newCompany.PIB,
		IdentificationNumber:         newCompany.IdentificationNumber,
		FirstAccountNumber:           newCompany.FirstAccountNumber,
		SecondAccountNumber:          newCompany.SecondAccountNumber,
		ForeignExchangeAccountNumber: newCompany.ForeignExchangeAccountNumber,
		CallNumber:                   newCompany.CallNumber,
		DateOfRegistration:           newCompany.DateOfRegistration,
		City:                         newCompany.City,
		ActivityCodeID:               newCompany.ActivityCodeID,
		MunicipalityID:               newCompany.MunicipalityID,
		EmployedByOtherFirm:          newCompany.EmployedByOtherFirm,
		EmploymentChanged:            newCompany.EmploymentChanged,
		Logo:                         newCompany.Logo,
		StreetAddress:                newCompany.StreetAddress,
		StreetNumber:                 newCompany.StreetNumber,
		Phone:                        newCompany.Phone,
		AgencyEmail:                  newCompany.AgencyEmail,
		SWIFT:                        newCompany.SWIFT,
		IBAN:                         newCompany.IBAN,
		Signature:                    newCompany.Signature,
		EmploymentType:               newCompany.EmploymentType,
		InvoiceDescription:           newCompany.InvoiceDescription,
	}
}

type CompanyService interface {
	GetAll(context.Context) (*[]Company, error)
	GetByID(ctx context.Context, id string) (*Company, error)
	Create(ctx context.Context, newCompany *NewCompany) (*Company, error)
	Update(ctx context.Context, id string, companyMap map[string]interface{}) (*Company, error)
	Delete(ctx context.Context, id string) error
}
