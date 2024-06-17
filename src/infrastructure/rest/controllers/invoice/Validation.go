package invoice

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateCreateInvoiceRequest validates CreateInvoiceRequest
func ValidateCreateInvoiceRequest(request *CreateInvoiceRequest) error {
	return validate.Struct(request)
}

// ValidateUpdateInvoiceRequest validates UpdateInvoiceRequest
func ValidateUpdateInvoiceRequest(request *UpdateInvoiceRequest) error {
	return validate.Struct(request)
}
