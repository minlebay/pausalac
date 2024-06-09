package user

import (
	domainUser "github.com/minlebay/pausalac/src/domain/user"
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
		ID:        user.ID.Hex(),
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
