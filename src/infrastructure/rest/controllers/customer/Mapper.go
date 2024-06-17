package customer

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pausalac/src/domain"
	"time"
)

// ToDomainMapper maps NewCustomer to Customer
func ToDomainMapper(newCustomer *domain.NewCustomer) *domain.Customer {
	return &domain.Customer{
		Name:               newCustomer.Name,
		TaxNumber:          newCustomer.TaxNumber,
		RegistrationNumber: newCustomer.RegistrationNumber,
		PhoneNumber:        newCustomer.PhoneNumber,
		Email:              newCustomer.Email,
		Address:            newCustomer.Address,
		City:               newCustomer.City,
		Country:            newCustomer.Country,
		Currency:           newCustomer.Currency,
		CustomerType:       newCustomer.CustomerType,
	}
}

// ToResponse maps Customer to CustomerResponse
func ToResponse(customer *domain.Customer) *CustomerResponse {
	return &CustomerResponse{
		ID:                 customer.Id.Hex(),
		Name:               customer.Name,
		TaxNumber:          customer.TaxNumber,
		RegistrationNumber: customer.RegistrationNumber,
		PhoneNumber:        customer.PhoneNumber,
		Email:              customer.Email,
		Address:            customer.Address,
		City:               customer.City,
		Country:            customer.Country,
		Currency:           customer.Currency,
		CustomerType:       string(customer.CustomerType),
	}
}

// ToDomainArray maps an array of Customers to an array of CustomerResponses
func ToResponseArray(customers *[]domain.Customer) *[]CustomerResponse {
	var responseCustomers []CustomerResponse
	for _, customer := range *customers {
		responseCustomers = append(responseCustomers, *ToResponse(&customer))
	}
	return &responseCustomers
}

// ToDomain converts CreateCustomerRequest to domain Customer
func ToDomain(req *CreateCustomerRequest) *domain.NewCustomer {
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
func ToDomainUpdate(req *UpdateCustomerRequest) map[string]interface{} {
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
