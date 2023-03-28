// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	authService "skyshi_gethired/application/usecases/auth"
	userRepository "skyshi_gethired/infrastructure/repository/postgres/user"
	authController "skyshi_gethired/infrastructure/restapi/controllers/auth"

	"gorm.io/gorm"
)

// AuthAdapter is a function that returns a auth controller
func AuthAdapter(db *gorm.DB) *authController.Controller {
	uRepository := userRepository.Repository{DB: db}
	service := authService.Service{UserRepository: uRepository}
	return &authController.Controller{AuthService: service}
}
