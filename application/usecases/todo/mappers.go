package todo

import (
	todoDomain "skyshi_gethired/domain/todo"
)

func (n *NewTodo) toDomainMapper(todoPriority string) *todoDomain.Todo {
	return &todoDomain.Todo{
		ActivityGroupID: *n.ActivityGroupID,
		Title:           *n.Title,
		IsActive:        *n.IsActive,
		Priority:        "very-high",
	}
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
