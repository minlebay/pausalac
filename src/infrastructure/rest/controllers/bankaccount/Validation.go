package bankaccount

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateCreateBankAccountRequest validates CreateBankAccountRequest
func ValidateCreateBankAccountRequest(request *CreateBankAccountRequest) error {
	return validate.Struct(request)
}

// ValidateUpdateBankAccountRequest validates UpdateBankAccountRequest
func ValidateUpdateBankAccountRequest(request *UpdateBankAccountRequest) error {
	return validate.Struct(request)
}
