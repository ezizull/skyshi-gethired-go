package adapter

import (
	activityService "skyshi_gethired/application/usecases/activity"
	activityRepository "skyshi_gethired/infrastructure/repository/mysql/activity"
	activityController "skyshi_gethired/infrastructure/restapi/controllers/activity"

	"gorm.io/gorm"
)

// ActivityAdapter is a function that returns a activity controller
func ActivityAdapter(db *gorm.DB) *activityController.Controller {
	mRepository := activityRepository.Repository{DB: db}
	service := activityService.Service{ActivityRepository: mRepository}
	return &activityController.Controller{ActivityService: service}
}
