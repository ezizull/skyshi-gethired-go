package todo

import (
	domainTodo "skyshi_gethired/domain/todo"
)

func (n *NewTodo) toDomainMapper() *domainTodo.Todo {
	return &domainTodo.Todo{
		ActivityGroupID: n.ActivityGroupID,
		Title:           n.Title,
		IsActive:        n.IsActive,
	}
}
