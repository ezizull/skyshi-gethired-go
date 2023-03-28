package todo

import domainTodo "skyshi_gethired/domain/todo"

func (todo *Todo) toDomainMapper() *domainTodo.Todo {
	return &domainTodo.Todo{
		ID:              todo.ID,
		Title:           todo.Title,
		ActivityGroupID: todo.ActivityGroupID,
		IsActive:        todo.IsActive,
		Priority:        todo.Priority,
		CreatedAt:       todo.CreatedAt,
		UpdatedAt:       todo.UpdatedAt,
	}
}

func fromDomainMapper(todo *domainTodo.Todo) *Todo {
	return &Todo{
		ID:              todo.ID,
		Title:           todo.Title,
		ActivityGroupID: todo.ActivityGroupID,
		IsActive:        todo.IsActive,
		Priority:        todo.Priority,
		CreatedAt:       todo.CreatedAt,
		UpdatedAt:       todo.UpdatedAt,
	}
}

func arrayToDomainMapper(todos *[]Todo) *[]domainTodo.Todo {
	todosDomain := make([]domainTodo.Todo, len(*todos))
	for i, todo := range *todos {
		todosDomain[i] = *todo.toDomainMapper()
	}

	return &todosDomain
}
