package user

// CreateUserRequest defines the request payload for creating a user
type CreateUserRequest struct {
	Author    string `json:"-"`
	UserName  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role" binding:"required"`
	Status    bool   `json:"status" binding:"required"`
}

// UpdateUserRequest defines the request payload for updating a user
type UpdateUserRequest struct {
	UserName  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string `json:"password,omitempty"`
	Role      string `json:"role,omitempty"`
	Status    bool   `json:"status,omitempty"`
}
