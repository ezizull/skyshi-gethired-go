package activity

import (
	activityUseCase "skyshi_gethired/application/usecases/activity"

	"github.com/gin-gonic/gin"
)

func createValidation(ctx *gin.Context) (activityBody *activityUseCase.NewActivity, message string) {
	// Get body data for newactivity
	_ = ctx.BindJSON(&activityBody)

	if activityBody.Title == nil {
		return activityBody, "title"
	}

	if activityBody.Email == nil {
		return activityBody, "email"
	}

	return activityBody, message
}

func updateValidation(ctx *gin.Context) (activityBody *activityUseCase.UpdateActivity, message string) {
	// Get body data for newactivity
	_ = ctx.BindJSON(&activityBody)

	if activityBody.Title == nil {
		return activityBody, "title"
	}

	return activityBody, message
}
