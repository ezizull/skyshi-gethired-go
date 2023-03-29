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

func (n *UpdateTodo) toDomainMapper() *todoDomain.Todo {
	return &todoDomain.Todo{
		Title:    *n.Title,
		IsActive: *n.IsActive,
		Priority: *n.Priority,
	}
}
