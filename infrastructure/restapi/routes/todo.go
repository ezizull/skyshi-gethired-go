// Package routes contains all routes of the application
package routes

import (
	todoController "skyshi_gethired/infrastructure/restapi/controllers/todo"

	"github.com/gin-gonic/gin"
)

// UserRoutes is a function that contains all routes of the todo
func TodoRoutes(router *gin.Engine, controller *todoController.Controller) {
	routerTodo := router.Group("/todo-items")
	{
		routerTodo.GET("", controller.GetAllTodos)
		routerTodo.GET("/:id", controller.GetTodoByID)
		routerTodo.POST("", controller.NewTodo)
		routerTodo.PATCH("/:id", controller.UpdateTodo)
		routerTodo.DELETE("/:id", controller.DeleteTodo)
	}
}
