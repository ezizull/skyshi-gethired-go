// Package todo contains the todo controller
package todo

import (
	useCaseTodo "skyshi_gethired/application/usecases/todo"

	"github.com/gin-gonic/gin"
)

func createValidation(ctx *gin.Context) (todoBody *useCaseTodo.NewTodo, message string) {
	// Get body data for newtodo
	_ = ctx.BindJSON(&todoBody)

	if todoBody.Title == nil {
		return nil, "title"
	}

	if todoBody.ActivityGroupID == nil || *todoBody.ActivityGroupID == 0 {
		return nil, "activity_group_id"
	}

	if todoBody.IsActive == nil {
		return nil, "is_active"
	}

	return todoBody, message
}
