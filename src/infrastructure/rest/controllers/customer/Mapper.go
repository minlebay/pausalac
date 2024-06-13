package customer

import (
	domainCustomer "pausalac/src/domain"
)

// ToDomainMapper maps NewCustomer to Customer
func ToDomainMapper(newCustomer *domainCustomer.NewCustomer) *domainCustomer.Customer {
	return &domainCustomer.Customer{
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
func ToResponse(customer *domainCustomer.Customer) *CustomerResponse {
	return &CustomerResponse{
		ID:                 customer.ID.Hex(),
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
func ToDomainArray(customers *[]domainCustomer.Customer) *[]CustomerResponse {
	var responseCustomers []CustomerResponse
	for _, customer := range *customers {
		responseCustomers = append(responseCustomers, *ToResponse(&customer))
	}
	return &responseCustomers
}
