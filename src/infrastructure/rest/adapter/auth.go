// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	authService "go_gin_api_clean/src/application/usecases/auth"
	userRepository "go_gin_api_clean/src/infrastructure/repository/user"
	authController "go_gin_api_clean/src/infrastructure/rest/controllers/auth"
	"gorm.io/gorm"
)

// AuthAdapter is a function that returns a auth controller
func AuthAdapter(db *gorm.DB) *authController.Controller {
	uRepository := userRepository.Repository{DB: db}
	service := authService.Service{UserRepository: uRepository}
	return &authController.Controller{AuthService: service}
}
