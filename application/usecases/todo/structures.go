package todo

import (
	domainTodo "skyshi_gethired/domain/todo"
)

// NewTodo is a struct that contains the data for a new todo
type NewTodo struct {
	ActivityGroupID uint   `json:"activity_group_id" example:"1" gorm:"foreignKey:ID"`
	Title           string `json:"title" example:"title todo"`
	IsActive        bool   `json:"is_active" example:"false"`
}

// UpdateTodo is a struct that contains the data for a update todo
type UpdateTodo struct {
	Title    string `json:"title,omitempty" example:"title todo"`
	Priority string `json:"priority,omitempty" example:"very-high"`
	IsActive bool   `json:"is_active,omitempty" example:"false"`
}

// PaginationResultTodo is a struct that contains the pagination result for todo
type PaginationResultTodo struct {
	Data       *[]domainTodo.Todo
	Total      int64
	Limit      int64
	Current    int64
	NextCursor uint
	PrevCursor uint
	NumPages   int64
}
