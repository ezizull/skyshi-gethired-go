package todo

import (
	activityDomain "skyshi_gethired/domain/activity"
	todoDomain "skyshi_gethired/domain/todo"
	todoRepository "skyshi_gethired/infrastructure/repository/mysql/todo"
)

// Service is a struct that contains the repository implementation for todo use case
type Service struct {
	TodoRepository todoRepository.Repository
}

// GetAll is a function that returns all todos
func (s *Service) GetAll() (todos []todoDomain.Todo, err error) {
	todos, err = s.TodoRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return todos, nil
}

// GetByActivity is a function that returns all todos
func (s *Service) GetByActivity(activityID string) (todos []todoDomain.Todo, err error) {
	todos, err = s.TodoRepository.GetByActivity(activityID)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

// GetByID is a function that returns a todo by id
func (s *Service) GetByID(id uint) (*todoDomain.Todo, error) {
	return s.TodoRepository.GetByID(id)
}

// GetActivity is a function that returns a todo by id
func (s *Service) GetActivity(id uint) (*activityDomain.Activity, error) {
	return s.TodoRepository.GetActivity(id)
}

// Create is a function that creates a todo
func (s *Service) Create(todo *NewTodo) (*todoDomain.Todo, error) {
	_, err := s.TodoRepository.GetActivity(*todo.ActivityGroupID)
	if err != nil {
		return nil, err
	}

	todoModel := todo.toDomainMapper("very-high")
	return s.TodoRepository.Create(todoModel)
}

// Update is a function that updates a todo by id
func (s *Service) Update(id uint, todo *UpdateTodo) (*todoDomain.Todo, error) {
	todoModel := todo.toDomainMapper()
	_, err := s.TodoRepository.Update(id, todoModel)
	if err != nil {
		return nil, err
	}

	return s.TodoRepository.GetByID(id)
}

// Delete is a function that deletes a todo by id
func (s *Service) Delete(id uint) error {
	_, err := s.TodoRepository.GetByID(id)
	if err != nil {
		return err
	}

	return s.TodoRepository.Delete(id)
}
