package todo

import (
	"net/http"
	todoUseCase "skyshi_gethired/application/usecases/todo"
	"skyshi_gethired/infrastructure/restapi/controllers"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	TodoService todoUseCase.Service
}

// GetAllTodos is the controller to getall todo
func (c *Controller) GetAllTodos(ctx *gin.Context) {
	activityGroupID := ctx.DefaultQuery("activity_group_id", "0")
	if activityGroupID != "0" {
		todos, err := c.TodoService.GetByActivity(activityGroupID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
				Status:  "Error",
				Message: err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, controllers.DefaultResponse{
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

	ctx.JSON(http.StatusOK, controllers.DefaultResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todos,
	})
}

// GetTodoByID is the controller to get a todo by id
func (c *Controller) GetTodoByID(ctx *gin.Context) {
	todoID := ctx.Param("id")
	todos, err := c.TodoService.GetByID(todoID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, controllers.ErrorResponse{
			Status:  "Not Found",
			Message: ("Todo with ID " + todoID + " Not Found"),
		})
		return
	}

	ctx.JSON(http.StatusOK, controllers.DefaultResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todos,
	})
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
		ctx.JSON(http.StatusNotFound, controllers.ErrorResponse{
			Status:  "Not Found",
			Message: ("Activity Group with ID " + strconv.Itoa(int(*todoBody.ActivityGroupID)) + " Not Found"),
		})
		return
	}

	ctx.JSON(http.StatusCreated, controllers.DefaultResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todos,
	})
}

// UpdateTodo is the controller to update a todo
func (c *Controller) UpdateTodo(ctx *gin.Context) {
	todoID := ctx.Param("id")

	// Get body data for updatetodo
	todoBody, message := updateValidation(ctx)
	if message != "" || todoBody == nil {
		ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
			Status:  "Bad Request",
			Message: (message + " cannot be null"),
		})
		return
	}

	// Get single todo for
	todos, err := c.TodoService.Update(todoID, todoBody)
	if err != nil {
		ctx.JSON(http.StatusNotFound, controllers.ErrorResponse{
			Status:  "Not Found",
			Message: ("Todo with ID " + todoID + " Not Found"),
		})
		return
	}

	ctx.JSON(http.StatusOK, controllers.DefaultResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todos,
	})
}

// DeleteTodo is the controller to delete a todo
func (c *Controller) DeleteTodo(ctx *gin.Context) {
	todoID := ctx.Param("id")
	err := c.TodoService.Delete(todoID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, controllers.ErrorResponse{
			Status:  "Not Found",
			Message: ("Todo with ID " + todoID + " Not Found"),
		})
		return
	}

	ctx.JSON(http.StatusOK, controllers.DefaultResponse{
		Status:  "Success",
		Message: "Success",
		Data:    gin.H{},
	})
}
