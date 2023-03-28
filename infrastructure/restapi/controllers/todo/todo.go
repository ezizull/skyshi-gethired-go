package todo

import (
	"errors"
	"net/http"
	useCaseTodo "skyshi_gethired/application/usecases/todo"
	domainError "skyshi_gethired/domain/errors"
	domainTodo "skyshi_gethired/domain/todo"
	"skyshi_gethired/infrastructure/restapi/controllers"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	TodoService useCaseTodo.Service
}

// NewTodo is the controller to create a todo
func (c *Controller) NewTodo(ctx *gin.Context) {
	var request NewTodoRequest

	if err := controllers.BindJSON(ctx, &request); err != nil {
		appError := domainError.NewAppError(err, domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	newTodo := useCaseTodo.NewTodo{
		Title:           request.Title,
		ActivityGroupID: request.ActivityGroupID,
		IsActive:        request.IsActive,
	}

	domainTodo, err := c.TodoService.Create(&newTodo)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, domainTodo)
}

// GetAllTodos is the controller to getall todo
func (c *Controller) GetAllTodos(ctx *gin.Context) {
	activityGroupIDStr := ctx.DefaultQuery("activity_group_id", "0")
	if activityGroupIDStr != "0" {
		todos, err := c.TodoService.GetByActivity(activityGroupIDStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
				Status:  "Error",
				Message: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusAccepted, controllers.DefaultResponse{
			Status:  "Success",
			Message: "Success",
			Data:    todos,
		})
		return
	}

	todos, err := c.TodoService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
			Status:  "Error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, controllers.DefaultResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todos,
	})
}

// GetTodoByID is the controller to get a todo by id
func (c *Controller) GetTodoByID(ctx *gin.Context) {
	todoID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := domainError.NewAppError(errors.New("todo id is invalid"), domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	domainTodo, err := c.TodoService.GetByID(todoID)
	if err != nil {
		appError := domainError.NewAppError(err, domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	ctx.JSON(http.StatusOK, domainTodo)
}

// UpdateTodo is the controller to update a todo
func (c *Controller) UpdateTodo(ctx *gin.Context) {
	todoID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		appError := domainError.NewAppError(errors.New("param id is necessary in the url"), domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	var requestMap map[string]interface{}

	err = controllers.BindJSONMap(ctx, &requestMap)
	if err != nil {
		appError := domainError.NewAppError(err, domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	var todo *domainTodo.Todo
	todo, err = c.TodoService.Update(uint(todoID), requestMap)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, todo)

}

// DeleteTodo is the controller to delete a todo
func (c *Controller) DeleteTodo(ctx *gin.Context) {
	todoID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		appError := domainError.NewAppError(errors.New("param id is necessary in the url"), domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	err = c.TodoService.Delete(todoID)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "resource deleted successfully"})
}
