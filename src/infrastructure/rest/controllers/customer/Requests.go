package customer

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
