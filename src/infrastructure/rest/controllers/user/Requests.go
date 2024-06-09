package user

import (
	domainUser "github.com/minlebay/pausalac/src/domain/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// CreateUserRequest defines the request payload for creating a user
type CreateUserRequest struct {
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

// ToDomain maps CreateUserRequest to NewUser
func (r *CreateUserRequest) ToDomain() *domainUser.NewUser {
	return &domainUser.NewUser{
		UserName:  r.UserName,
		Email:     r.Email,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Password:  r.Password,
		Role:      r.Role,
		Status:    r.Status,
	}
}

// ToDomainUpdate maps UpdateUserRequest to map[string]interface{}
func (r *UpdateUserRequest) ToDomainUpdate() map[string]interface{} {
	updateMap := make(map[string]interface{})
	if r.UserName != "" {
		updateMap["username"] = r.UserName
	}
	if r.Email != "" {
		updateMap["email"] = r.Email
	}
	if r.FirstName != "" {
		updateMap["first_name"] = r.FirstName
	}
	if r.LastName != "" {
		updateMap["last_name"] = r.LastName
	}
	if r.Password != "" {
		updateMap["hash_password"] = r.Password // Обработка хэширования пароля
	}
	if r.Role != "" {
		updateMap["role"] = r.Role
	}
	updateMap["status"] = r.Status
	updateMap["updated_at"] = primitive.NewDateTimeFromTime(time.Now())
	return updateMap
}
