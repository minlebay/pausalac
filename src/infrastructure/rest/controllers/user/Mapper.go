package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	domainUser "pausalac/src/domain"
	"time"
)

// ToDomainMapper maps NewUser to User
func ToDomainMapper(newUser *domainUser.NewUser) *domainUser.User {
	return &domainUser.User{
		UserName:     newUser.UserName,
		Email:        newUser.Email,
		FirstName:    newUser.FirstName,
		LastName:     newUser.LastName,
		Role:         newUser.Role,
		Status:       newUser.Status,
		HashPassword: newUser.Password, // Обработка хэширования пароля
	}
}

// ToResponse maps User to UserResponse
func ToResponse(user *domainUser.User) *UserResponse {
	return &UserResponse{
		Id:        user.Id.Hex(),
		UserName:  user.UserName,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Status:    user.Status,
		Role:      user.Role,
	}
}

// ToDomainArray maps array of Users to array of UserResponses
func ToDomainArray(users *[]domainUser.User) *[]UserResponse {
	var responseUsers []UserResponse
	for _, user := range *users {
		responseUsers = append(responseUsers, *ToResponse(&user))
	}
	return &responseUsers
}

// ToDomain maps CreateUserRequest to NewUser
func ToDomain(r *CreateUserRequest) *domainUser.NewUser {
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
func ToDomainUpdate(r *UpdateUserRequest) map[string]interface{} {
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
		updateMap["hash_password"] = r.Password
	}
	if r.Role != "" {
		updateMap["role"] = r.Role
	}
	updateMap["status"] = r.Status
	updateMap["updated_at"] = primitive.NewDateTimeFromTime(time.Now())
	return updateMap
}
