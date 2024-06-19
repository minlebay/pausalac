package company

import (
	"pausalac/src/infrastructure/rest/controllers/bankaccount"
	"time"
)

// CompanyResponse defines the response payload for a company
type CompanyResponse struct {
	Id                           string                            `json:"id"`
	Author                       string                            `json:"user_id"`
	AgencyId                     string                            `json:"agency_id"`
	Name                         string                            `json:"name"`
	FullName                     string                            `json:"full_name"`
	PIB                          string                            `json:"pib"`
	IdentificationNumber         string                            `json:"identification_number"`
	FirstAccountNumber           string                            `json:"first_account_number"`
	SecondAccountNumber          string                            `json:"second_account_number"`
	ForeignExchangeAccountNumber string                            `json:"foreign_exchange_account_number"`
	CallNumber                   string                            `json:"call_number"`
	DateOfRegistration           string                            `json:"date_of_registration"`
	City                         string                            `json:"city"`
	ActivityCodeId               string                            `json:"activity_code_id"`
	MunicipalityId               string                            `json:"municipality_id"`
	EmployedByOtherFirm          string                            `json:"employed_by_other_firm"`
	EmploymentChanged            string                            `json:"employment_changed"`
	Logo                         string                            `json:"logo"`
	StreetAddress                string                            `json:"street_address"`
	StreetNumber                 string                            `json:"street_number"`
	Phone                        string                            `json:"phone"`
	AgencyEmail                  string                            `json:"agency_email"`
	SWIFT                        string                            `json:"swift"`
	IBAN                         string                            `json:"iban"`
	Signature                    string                            `json:"signature"`
	EmploymentType               string                            `json:"employment_type"`
	InvoiceDescription           string                            `json:"invoice_description"`
	CreatedAt                    time.Time                         `json:"created_at"`
	UpdatedAt                    time.Time                         `json:"updated_at"`
	BankAccounts                 []bankaccount.BankAccountResponse `json:"bank_accounts"`
}

// MessageResponse defines a generic response message
type MessageResponse struct {
	Message string `json:"message"`
}
