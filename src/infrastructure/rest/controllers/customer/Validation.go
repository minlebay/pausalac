package customer

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateCreateCustomerRequest validates CreateCustomerRequest
func ValidateCreateCustomerRequest(request *CreateCustomerRequest) error {
	return validate.Struct(request)
}

// ValidateUpdateCustomerRequest validates UpdateCustomerRequest
func ValidateUpdateCustomerRequest(request *UpdateCustomerRequest) error {
	return validate.Struct(request)
}
