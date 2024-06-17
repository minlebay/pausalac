package invoice

import (
	"pausalac/src/infrastructure/rest/controllers/bankaccount"
	"pausalac/src/infrastructure/rest/controllers/customer"
	"pausalac/src/infrastructure/rest/controllers/service"
	"pausalac/src/infrastructure/rest/controllers/user"
)

type CreateInvoiceRequest struct {
	Author        user.CreateUserRequest               `json:"author" binding:"required"`
	Created       string                               `json:"created" binding:"required"`
	BankAccount   bankaccount.CreateBankAccountRequest `json:"bank_account"`
	Cancelled     string                               `json:"cancelled"`
	Client        customer.CreateCustomerRequest       `json:"client" binding:"required"`
	Comment       string                               `json:"comment"`
	Currency      string                               `json:"currency"`
	IBAN          string                               `json:"iban"`
	Date          string                               `json:"date"`
	Number        string                               `json:"number"`
	PaidDate      string                               `json:"paid_date"`
	PaidValue     string                               `json:"paid_value"`
	SentDate      string                               `json:"sent_date"`
	Services      service.CreateServiceArrayRequest    `json:"services"`
	Status        string                               `json:"status"`
	SWIFT         string                               `json:"swift"`
	TradingDate   string                               `json:"trading_date"`
	TraidingPlace string                               `json:"traiding_place"`
	Type          string                               `json:"type"`
	ValueInRSD    string                               `json:"value_in_rsd"`
}

type UpdateInvoiceRequest struct {
	Author        user.UpdateUserRequest               `json:"author" binding:"required"`
	Created       string                               `json:"created" binding:"required"`
	BankAccount   bankaccount.UpdateBankAccountRequest `json:"bank_account"`
	Cancelled     string                               `json:"cancelled"`
	Client        customer.UpdateCustomerRequest       `json:"client" binding:"required"`
	Comment       string                               `json:"comment"`
	Currency      string                               `json:"currency"`
	IBAN          string                               `json:"iban"`
	Date          string                               `json:"date"`
	Number        string                               `json:"number"`
	PaidDate      string                               `json:"paid_date"`
	PaidValue     string                               `json:"paid_value"`
	SentDate      string                               `json:"sent_date"`
	Services      service.UpdateServiceArrayRequest    `json:"services"`
	Status        string                               `json:"status"`
	SWIFT         string                               `json:"swift"`
	TradingDate   string                               `json:"trading_date"`
	TraidingPlace string                               `json:"traiding_place"`
	Type          string                               `json:"type"`
	ValueInRSD    string                               `json:"value_in_rsd"`
}
