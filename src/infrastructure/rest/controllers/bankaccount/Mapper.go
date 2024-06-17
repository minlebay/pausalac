package bankaccount

import (
	"pausalac/src/domain"
	"time"
)

// ToResponse maps BankAccount to BankAccountResponse
func ToResponse(bankAccount *domain.BankAccount) BankAccountResponse {
	return BankAccountResponse{
		Id:            bankAccount.Id.Hex(),
		UserID:        bankAccount.UserID,
		AccountNumber: bankAccount.AccountNumber,
		BankName:      bankAccount.BankName,
		SwiftCode:     bankAccount.SwiftCode,
		IBAN:          bankAccount.IBAN,
		Currency:      bankAccount.Currency,
		CreatedAt:     bankAccount.CreatedAt,
		UpdatedAt:     bankAccount.UpdatedAt,
	}
}

// ToDomainArray maps an array of BankAccount to BankAccountResponse array
func ToResponseArray(bankAccounts *[]domain.BankAccount) []BankAccountResponse {
	var bankAccountResponses []BankAccountResponse
	for _, bankAccount := range *bankAccounts {
		bankAccountResponses = append(bankAccountResponses, ToResponse(&bankAccount))
	}
	return bankAccountResponses
}

func ToDomain(req *CreateBankAccountRequest) *domain.NewBankAccount {
	return &domain.NewBankAccount{
		UserID:        req.UserID,
		AccountNumber: req.AccountNumber,
		BankName:      req.BankName,
		SwiftCode:     req.SwiftCode,
		IBAN:          req.IBAN,
		Currency:      req.Currency,
	}
}

func ToDomainUpdate(req *UpdateBankAccountRequest) map[string]interface{} {
	bankAccountMap := make(map[string]interface{})
	if req.UserID != "" {
		bankAccountMap["user_id"] = req.UserID
	}
	if req.AccountNumber != "" {
		bankAccountMap["account_number"] = req.AccountNumber
	}
	if req.BankName != "" {
		bankAccountMap["bank_name"] = req.BankName
	}
	if req.SwiftCode != "" {
		bankAccountMap["swift_code"] = req.SwiftCode
	}
	if req.IBAN != "" {
		bankAccountMap["iban"] = req.IBAN
	}
	if req.Currency != "" {
		bankAccountMap["currency"] = req.Currency
	}
	bankAccountMap["updated_at"] = time.Now()
	return bankAccountMap
}
