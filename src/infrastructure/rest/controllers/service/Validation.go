package service

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateCreateServiceRequest validates CreateServiceRequest
func ValidateCreateServiceRequest(request *CreateServiceRequest) error {
	return validate.Struct(request)
}

// ValidateUpdateServiceRequest validates UpdateServiceRequest
func ValidateUpdateServiceRequest(request *UpdateServiceRequest) error {
	return validate.Struct(request)
}
