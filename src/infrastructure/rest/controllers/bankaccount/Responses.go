package bankaccount

import "time"

// BankAccountResponse defines the response payload for a bank account
type BankAccountResponse struct {
	Id            string    `json:"id"`
	UserID        string    `json:"user_id"`
	AccountNumber string    `json:"account_number"`
	BankName      string    `json:"bank_name"`
	SwiftCode     string    `json:"swift_code"`
	IBAN          string    `json:"iban"`
	Currency      string    `json:"currency"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// MessageResponse defines a generic response message
type MessageResponse struct {
	Message string `json:"message"`
}
