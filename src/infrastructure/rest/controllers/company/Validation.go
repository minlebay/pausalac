package company

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateCreateCompanyRequest validates CreateCompanyRequest
func ValidateCreateCompanyRequest(request *CreateCompanyRequest) error {
	return validate.Struct(request)
}

// ValidateUpdateCompanyRequest validates UpdateCompanyRequest
func ValidateUpdateCompanyRequest(request *UpdateCompanyRequest) error {
	return validate.Struct(request)
}
