package customer

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pausalac/src/domain"
	"time"
)

// CreateCustomerRequest defines the request payload for creating a customer
type CreateCustomerRequest struct {
	Name               string `json:"name" binding:"required"`
	TaxNumber          string `json:"tax_number" binding:"required"`
	RegistrationNumber string `json:"registration_number" binding:"required"`
	PhoneNumber        string `json:"phone_number" binding:"required"`
	Email              string `json:"email" binding:"required,email"`
	Address            string `json:"address" binding:"required"`
	City               string `json:"city" binding:"required"`
	Country            string `json:"country" binding:"required"`
	Currency           string `json:"currency" binding:"required"`
	CustomerType       string `json:"customer_type" binding:"required"`
}

// UpdateCustomerRequest defines the request payload for updating a customer
type UpdateCustomerRequest struct {
	Name               string `json:"name,omitempty"`
	TaxNumber          string `json:"tax_number,omitempty"`
	RegistrationNumber string `json:"registration_number,omitempty"`
	PhoneNumber        string `json:"phone_number,omitempty"`
	Email              string `json:"email,omitempty"`
	Address            string `json:"address,omitempty"`
	City               string `json:"city,omitempty"`
	Country            string `json:"country,omitempty"`
	Currency           string `json:"currency,omitempty"`
	CustomerType       string `json:"customer_type,omitempty"`
}

// ToDomain converts CreateCustomerRequest to domain Customer
func (req *CreateCustomerRequest) ToDomain() *domain.NewCustomer {
	return &domain.NewCustomer{
		Name:               req.Name,
		TaxNumber:          req.TaxNumber,
		RegistrationNumber: req.RegistrationNumber,
		PhoneNumber:        req.PhoneNumber,
		Email:              req.Email,
		Address:            req.Address,
		City:               req.City,
		Country:            req.Country,
		Currency:           req.Currency,
		CustomerType:       domain.CustomerType(req.CustomerType),
	}
}

// ToDomainUpdate converts UpdateCustomerRequest to domain Customer
func (req *UpdateCustomerRequest) ToDomainUpdate() map[string]interface{} {
	customerMap := make(map[string]interface{})
	if req.Name != "" {
		customerMap["name"] = req.Name
	}
	if req.TaxNumber != "" {
		customerMap["tax_number"] = req.TaxNumber
	}
	if req.RegistrationNumber != "" {
		customerMap["registration_number"] = req.RegistrationNumber
	}
	if req.PhoneNumber != "" {
		customerMap["phone_number"] = req.PhoneNumber
	}
	if req.Email != "" {
		customerMap["email"] = req.Email
	}
	if req.Address != "" {
		customerMap["address"] = req.Address
	}
	if req.City != "" {
		customerMap["city"] = req.City
	}
	if req.Country != "" {
		customerMap["country"] = req.Country
	}
	if req.Currency != "" {
		customerMap["currency"] = req.Currency
	}
	if req.CustomerType != "" {
		customerMap["customer_type"] = req.CustomerType
	}
	customerMap["updated_at"] = primitive.NewDateTimeFromTime(time.Now())
	return customerMap
}
