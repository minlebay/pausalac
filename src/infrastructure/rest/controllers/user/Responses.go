package user

// MessageResponse is a struct that contains the response body for the message
type MessageResponse struct {
	Message string `json:"message"`
}

// UserResponse defines the response payload for a user
type UserResponse struct {
	Id        string `json:"id"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Status    bool   `json:"status"`
	Role      string `json:"role"`
}
