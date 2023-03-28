package todo

import (
	domainTodo "skyshi_gethired/domain/todo"
	todoDomain "skyshi_gethired/domain/todo"
	todoRepository "skyshi_gethired/infrastructure/repository/mysql/todo"
)

// Service is a struct that contains the repository implementation for todo use case
type Service struct {
	TodoRepository todoRepository.Repository
}

// GetAll is a function that returns all medicines
func (s *Service) GetAll(page int64, limit int64) (todos *[]domainTodo.Todo, err error) {

	todos, err = s.TodoRepository.GetAll(page, limit)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

// GetByID is a function that returns a todo by id
func (s *Service) GetByID(id int) (*todoDomain.Todo, error) {
	return s.TodoRepository.GetByID(id)
}

// Create is a function that creates a todo
func (s *Service) Create(todo *NewTodo) (*todoDomain.Todo, error) {

	todoModel := todo.toDomainMapper()

	return s.TodoRepository.Create(todoModel)
}

// GetByMap is a function that returns a todo by map
func (s *Service) GetByMap(medicineMap map[string]interface{}) (*todoDomain.Todo, error) {
	return s.TodoRepository.GetOneByMap(medicineMap)
}

// Delete is a function that deletes a todo by id
func (s *Service) Delete(id int) error {
	return s.TodoRepository.Delete(id)
}

// Update is a function that updates a todo by id
func (s *Service) Update(id uint, medicineMap map[string]interface{}) (*todoDomain.Todo, error) {
	return s.TodoRepository.Update(id, medicineMap)
}
