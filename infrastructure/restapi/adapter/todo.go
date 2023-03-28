package adapter

import (
	todoService "skyshi_gethired/application/usecases/todo"
	todoRepository "skyshi_gethired/infrastructure/repository/mysql/todo"
	todoController "skyshi_gethired/infrastructure/restapi/controllers/todo"

	"gorm.io/gorm"
)

// TodoAdapter is a function that returns a todo controller
func TodoAdapter(db *gorm.DB) *todoController.Controller {
	mRepository := todoRepository.Repository{DB: db}
	service := todoService.Service{TodoRepository: mRepository}
	return &todoController.Controller{TodoService: service}
}
