package user

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateCreateUserRequest validates CreateUserRequest
func ValidateCreateUserRequest(request *CreateUserRequest) error {
	return validate.Struct(request)
}

// ValidateUpdateUserRequest validates UpdateUserRequest
func ValidateUpdateUserRequest(request *UpdateUserRequest) error {
	return validate.Struct(request)
}
