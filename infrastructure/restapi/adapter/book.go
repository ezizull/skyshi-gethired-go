// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	bookService "skyshi_gethired/application/usecases/book"
	bookRepository "skyshi_gethired/infrastructure/repository/postgres/book"
	bookController "skyshi_gethired/infrastructure/restapi/controllers/book"

	"gorm.io/gorm"
)

// BookAdapter is a function that returns a book controller
func BookAdapter(db *gorm.DB) *bookController.Controller {
	mRepository := bookRepository.Repository{DB: db}
	service := bookService.Service{BookRepository: mRepository}
	return &bookController.Controller{BookService: service}
}
