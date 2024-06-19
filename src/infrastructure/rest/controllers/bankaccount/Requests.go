package bankaccount

// CreateBankAccountRequest defines the request payload for creating a bank account
type CreateBankAccountRequest struct {
	Author        string `json:"-"`
	AccountNumber string `json:"account_number" binding:"required"`
	BankName      string `json:"bank_name" binding:"required"`
	SwiftCode     string `json:"swift_code" binding:"required"`
	IBAN          string `json:"iban" binding:"required"`
	Currency      string `json:"currency" binding:"required"`
}

// UpdateBankAccountRequest defines the request payload for updating a bank account
type UpdateBankAccountRequest struct {
	AccountNumber string `json:"account_number"`
	BankName      string `json:"bank_name"`
	SwiftCode     string `json:"swift_code"`
	IBAN          string `json:"iban"`
	Currency      string `json:"currency"`
}
