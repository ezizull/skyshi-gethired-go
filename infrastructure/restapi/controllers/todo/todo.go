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

// NewTodo godoc
// @Tags todo
// @Summary Create New Todo
// @Descriptioniption Create new todo on the system
// @Accept  json
// @Produce  json
// @Param data body NewTodoRequest true "body data"
// @Success 200 {object} domainTodo.Todo
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /todo [post]
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

// GetAllTodos godoc
// @Tags todo
// @Summary Get all Todos
// @Description Get all Todos on the system
// @Param   limit  query   string  true        "limit"
// @Param   page  query   string  true        "page"
// @Success 200 {object} []useCaseTodo.PaginationResultTodo
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /todo [get]
func (c *Controller) GetAllTodos(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "20")

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		appError := domainError.NewAppError(errors.New("param page is necessary to be an integer"), domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		appError := domainError.NewAppError(errors.New("param limit is necessary to be an integer"), domainError.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	todos, err := c.TodoService.GetAll(page, limit)
	if err != nil {
		appError := domainError.NewAppErrorWithType(domainError.UnknownError)
		_ = ctx.Error(appError)
		return
	}
	ctx.JSON(http.StatusOK, todos)
}

// GetTodoByID godoc
// @Tags todo
// @Summary Get todos by ID
// @Descriptioniption Get Todos by ID on the system
// @Param todo_id path int true "id of todo"
// @Success 200 {object} domainTodo.Todo
// @Failure 400 {object} MessageResponse
// @Failure 500 {object} MessageResponse
// @Router /todo/{todo_id} [get]
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
