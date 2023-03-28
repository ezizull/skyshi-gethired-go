package todo

import (
	domainTodo "skyshi_gethired/domain/todo"
	"time"

	"gorm.io/gorm"
)

// Todo is a struct that contains the todo model
type Todo struct {
	ID              uint      `json:"id" gorm:"primary_key"`
	ActivityGroupID uint      `json:"activity_group_id" example:"1"`
	Title           string    `json:"title" example:"title todo"`
	IsActive        bool      `json:"is_active" example:"false" gorm:"default:true"`
	Priority        string    `json:"priority" example:"very-high"`
	CreatedAt       time.Time `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt       time.Time `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
	DeletedAt       *gorm.DeletedAt
}

// TableName overrides the table name used by User to `users`
func (*Todo) TableName() string {
	return "todos"
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
