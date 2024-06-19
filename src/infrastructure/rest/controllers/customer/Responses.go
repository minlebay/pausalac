package customer

// CustomerResponse defines the response payload for a customer
type CustomerResponse struct {
	Id                 string `json:"id"`
	Author             string `json:"author"`
	Name               string `json:"name"`
	TaxNumber          string `json:"tax_number"`
	RegistrationNumber string `json:"registration_number"`
	PhoneNumber        string `json:"phone_number"`
	Email              string `json:"email"`
	Address            string `json:"address"`
	City               string `json:"city"`
	Country            string `json:"country"`
	Currency           string `json:"currency"`
	CustomerType       string `json:"customer_type"`
}

// MessageResponse defines a generic response message
type MessageResponse struct {
	Message string `json:"message"`
}
