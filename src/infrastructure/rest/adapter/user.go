// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	userService "go_gin_api_clean/src/application/usecases/user"
	userRepository "go_gin_api_clean/src/infrastructure/repository/user"
	userController "go_gin_api_clean/src/infrastructure/rest/controllers/user"
	"gorm.io/gorm"
)

// UserAdapter is a function that returns a user controller
func UserAdapter(db *gorm.DB) *userController.Controller {
	uRepository := userRepository.Repository{DB: db}
	service := userService.Service{UserRepository: uRepository}
	return &userController.Controller{UserService: service}
}
