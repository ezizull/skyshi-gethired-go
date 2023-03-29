// Package routes contains all routes of the application
package routes

import (
	activityController "skyshi_gethired/infrastructure/restapi/controllers/activity"

	"github.com/gin-gonic/gin"
)

// UserRoutes is a function that contains all routes of the activity
func ActivityRoutes(router *gin.Engine, controller *activityController.Controller) {
	routerActivity := router.Group("/activity-groups")
	{
		routerActivity.GET("", controller.GetActivities)
		routerActivity.GET("/:id", controller.GetActivityByID)
		routerActivity.POST("", controller.NewActivity)
		routerActivity.PATCH("/:id", controller.UpdateActivity)
		routerActivity.DELETE("/:id", controller.DeleteActivity)
	}
}
