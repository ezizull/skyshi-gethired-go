package todo

import (
	todoDomain "skyshi_gethired/domain/todo"
)

func (n *NewTodo) toDomainMapper(todoPriority string) (todo *todoDomain.Todo) {
	todo = &todoDomain.Todo{}

	if n.ActivityGroupID != nil {
		todo.ActivityGroupID = *n.ActivityGroupID
	}

	if n.Title != nil {
		todo.Title = *n.Title
	}

	if n.IsActive != nil {
		todo.IsActive = *n.IsActive
	} else {
		todo.IsActive = true
	}

	todo.Priority = "very-high"
	return todo
}

func (n *UpdateTodo) toDomainMapper() (todo *todoDomain.Todo) {
	todo = &todoDomain.Todo{}

	if n.Title != nil {
		todo.Title = *n.Title
	}

	if n.IsActive != nil {
		todo.IsActive = *n.IsActive
	}

	if n.Priority != nil {
		todo.Priority = *n.Priority
	}

	return todo
}
