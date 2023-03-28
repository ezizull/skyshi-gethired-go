// Package medicine contains the medicine controller
package todo

// NewTodoRequest is a struct that contains the new medicine request information
type NewTodoRequest struct {
	ActivityGroupID uint   `json:"activity_group_id" example:"1" binding:"required"`
	Title           string `json:"title" example:"title todo" binding:"required"`
	IsActive        bool   `json:"is_active" example:"false" binding:"required"`
}
