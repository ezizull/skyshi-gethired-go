package activity

import (
	"net/http"
	activityUseCase "skyshi_gethired/application/usecases/activity"
	"skyshi_gethired/infrastructure/restapi/controllers"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	ActivityService activityUseCase.Service
}

// GetActivities is the controller to getall activity
func (c *Controller) GetActivities(ctx *gin.Context) {
	activities, err := c.ActivityService.GetAll()
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
		Data:    activities,
	})
}

// GetActivityByID is the controller to get a activity by id
func (c *Controller) GetActivityByID(ctx *gin.Context) {
	activityID := ctx.Param("id")
	activities, err := c.ActivityService.GetByID(activityID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, controllers.ErrorResponse{
			Status:  "Not Found",
			Message: ("Activity with ID " + activityID + " Not Found"),
		})
		return
	}

	ctx.JSON(http.StatusOK, controllers.DefaultResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activities,
	})
}

// NewActivity is the controller to create a activity
func (c *Controller) NewActivity(ctx *gin.Context) {
	// validation create activity body
	activityBody, message := createValidation(ctx)
	if message != "" || activityBody == nil {
		ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
			Status:  "Bad Request",
			Message: (message + " cannot be null"),
		})
		return
	}

	activities, err := c.ActivityService.Create(activityBody)
	if err != nil {
		ctx.JSON(http.StatusNotFound, controllers.ErrorResponse{
			Status:  "Error",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, controllers.DefaultResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activities,
	})
}

// UpdateActivity is the controller to update a activity
func (c *Controller) UpdateActivity(ctx *gin.Context) {
	activityID := ctx.Param("id")

	// Get body data for updateactivity
	activityBody, message := updateValidation(ctx)
	if message != "" || activityBody == nil {
		ctx.JSON(http.StatusBadRequest, controllers.ErrorResponse{
			Status:  "Bad Request",
			Message: (message + " cannot be null"),
		})
		return
	}

	// Get single activity for
	activities, err := c.ActivityService.Update(activityID, activityBody)
	if err != nil || activities == nil {
		ctx.JSON(http.StatusNotFound, controllers.ErrorResponse{
			Status:  "Not Found",
			Message: ("Activity with ID " + activityID + " Not Found"),
		})
		return
	}

	ctx.JSON(http.StatusOK, controllers.DefaultResponse{
		Status:  "Success",
		Message: "Success",
		Data:    activities,
	})
}

// DeleteActivity is the controller to delete a activity
func (c *Controller) DeleteActivity(ctx *gin.Context) {
	activityID := ctx.Param("id")
	err := c.ActivityService.Delete(activityID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, controllers.ErrorResponse{
			Status:  "Not Found",
			Message: ("Activity with ID " + activityID + " Not Found"),
		})
		return
	}

	ctx.JSON(http.StatusOK, controllers.DefaultResponse{
		Status:  "Success",
		Message: "Success",
		Data:    gin.H{},
	})
}
