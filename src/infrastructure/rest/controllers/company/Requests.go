package company

import "pausalac/src/infrastructure/rest/controllers/bankaccount"

// CreateCompanyRequest defines the request payload for creating a company
type CreateCompanyRequest struct {
	Author                       string                                 `json:"-"`
	AgencyId                     string                                 `json:"agency_id"`
	Name                         string                                 `json:"name" binding:"required"`
	FullName                     string                                 `json:"full_name"`
	PIB                          string                                 `json:"pib"`
	IdentificationNumber         string                                 `json:"identification_number"`
	ForeignExchangeAccountNumber string                                 `json:"foreign_exchange_account_number"`
	CallNumber                   string                                 `json:"call_number"`
	DateOfRegistration           string                                 `json:"date_of_registration"`
	City                         string                                 `json:"city"`
	ActivityCodeId               string                                 `json:"activity_code_id"`
	MunicipalityId               string                                 `json:"municipality_id"`
	Logo                         string                                 `json:"logo"`
	StreetAddress                string                                 `json:"street_address"`
	StreetNumber                 string                                 `json:"street_number"`
	Phone                        string                                 `json:"phone"`
	AgencyEmail                  string                                 `json:"agency_email"`
	Signature                    string                                 `json:"signature"`
	EmploymentType               string                                 `json:"employment_type"`
	InvoiceDescription           string                                 `json:"invoice_description"`
	BankAccounts                 []bankaccount.CreateBankAccountRequest `json:"bank_accounts"`
}

// UpdateCompanyRequest defines the request payload for updating a company
type UpdateCompanyRequest struct {
	Author               string
	AgencyId             string                                 `json:"agency_id"`
	Name                 string                                 `json:"name"`
	FullName             string                                 `json:"full_name"`
	PIB                  string                                 `json:"pib"`
	IdentificationNumber string                                 `json:"identification_number"`
	FirstAccountNumber   string                                 `json:"first_account_number"`
	CallNumber           string                                 `json:"call_number"`
	DateOfRegistration   string                                 `json:"date_of_registration"`
	City                 string                                 `json:"city"`
	ActivityCodeId       string                                 `json:"activity_code_id"`
	MunicipalityId       string                                 `json:"municipality_id"`
	Logo                 string                                 `json:"logo"`
	StreetAddress        string                                 `json:"street_address"`
	StreetNumber         string                                 `json:"street_number"`
	Phone                string                                 `json:"phone"`
	AgencyEmail          string                                 `json:"agency_email"`
	Signature            string                                 `json:"signature"`
	EmploymentType       string                                 `json:"employment_type"`
	InvoiceDescription   string                                 `json:"invoice_description"`
	BankAccounts         []bankaccount.UpdateBankAccountRequest `json:"bank_accounts"`
}
