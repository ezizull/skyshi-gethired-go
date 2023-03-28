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
	// validation create todo body
	todoBody, message := createValidation(ctx)
	if message != "" || todoBody == nil {
		ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
			Status:  "Bad Request",
			Message: (message + " cannot be null"),
		})
		return
	}

	todos, err := c.TodoService.Create(todoBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
			Status:  "Error",
			Message: err.Error(),
		})
		return
	}

	// if todos.ID == 0 {
	// 	ctx.JSON(http.StatusNotFound, controllers.ErrorResponse{
	// 		Status:  "Not Found",
	// 		Message: ("Activity Group with ID " + strconv.Itoa(int(*todoBody.ActivityGroupID)) + " Not Found"),
	// 	})
	// 	return
	// }

	ctx.JSON(http.StatusAccepted, controllers.DefaultResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todos,
	})
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
	todoIDStr := ctx.Param("id")
	todoID, err := strconv.ParseInt(todoIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusNotFound, controllers.ErrorResponse{
			Status:  "Not Found",
			Message: ("Todo with ID " + todoIDStr + " Not Found"),
		})
		return
	}

	todos, err := c.TodoService.GetByID(int(todoID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, controllers.ErrorResponse{
			Status:  "Not Found",
			Message: ("Todo with ID " + todoIDStr + " Not Found"),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, controllers.DefaultResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todos,
	})
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
