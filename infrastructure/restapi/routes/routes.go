// Package routes contains all routes of the application
package routes

import (
	// swaggerFiles for documentation
	_ "skyshi_gethired/docs"
	"skyshi_gethired/infrastructure/restapi/adapter"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Security is a struct that contains the security of the application
// @SecurityDefinitions.jwt
type Security struct {
	Authorization string `header:"Authorization" json:"Authorization"`
}

func ApplicationV1Router(router *gin.Engine, db *gorm.DB) {
	// TodoRoutes is a routes group of the todo
	TodoRoutes(router, adapter.TodoAdapter(db))

	// ActivityRoutes is a routes group of the todo
	ActivityRoutes(router, adapter.ActivityAdapter(db))
}
