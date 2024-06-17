package invoice

import (
	"pausalac/src/domain"
)

type InvoiceResponse struct {
	Id            string      `json:"id"`
	Author        domain.User `json:"author"`
	Created       string      `json:"created"`
	BankAccount   string      `json:"bank_account"`
	Cancelled     bool        `json:"cancelled"`
	Client        string      `json:"client"`
	CreatedAt     string      `json:"created_at"`
	Comment       string      `json:"comment"`
	Currency      string      `json:"currency"`
	IBAN          string      `json:"iban"`
	Date          string      `json:"date"`
	Number        string      `json:"number"`
	PaidDate      string      `json:"paid_date"`
	PaidValue     string      `json:"paid_value"`
	SentDate      string      `json:"sent_date"`
	Services      string      `json:"services"`
	Status        string      `json:"status"`
	SWIFT         string      `json:"swift"`
	TradingDate   string      `json:"trading_date"`
	TraidingPlace string      `json:"traiding_place"`
	Type          string      `json:"type"`
	ValueInRSD    string      `json:"value_in_rsd"`
	UpdatedAt     string      `json:"updated_at"`
}

// MessageResponse defines a generic response message
type MessageResponse struct {
	Message string `json:"message"`
}
