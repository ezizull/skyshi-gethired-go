package activity

import (
	domainActivity "skyshi_gethired/domain/activity"

	"gorm.io/gorm"
)

// Repository is a struct that contains the database implementation for medicine entity
type Repository struct {
	DB *gorm.DB
}

// GetAll Fetch all activity data
func (r *Repository) GetAll() (activity []domainActivity.Activity, err error) {
	resp := r.DB.Find(&activity)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return activity, nil
}

// GetByActivity Fetch all activity data
func (r *Repository) GetByActivity(activityID string) (activity []domainActivity.Activity, err error) {
	resp := r.DB.Where("activity_group_id = ?", activityID).Find(&activity)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return activity, nil
}

// Create ... Insert New data
func (r *Repository) Create(newTodo *domainActivity.Activity) (*domainActivity.Activity, error) {
	resp := r.DB.Create(newTodo)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return newTodo, nil
}

// GetByID ... Fetch only one activity by Id
func (r *Repository) GetByID(id int) (activity *domainActivity.Activity, err error) {
	err = r.DB.Where("id = ?", id).First(&activity).Error
	if err != nil {
		return nil, err
	}

	return activity, nil
}

// Update ... Update activity
func (r *Repository) Update(id uint, todoMap map[string]interface{}) (*domainActivity.Activity, error) {
	return nil, nil
}

// Delete ... Delete activity
func (r *Repository) Delete(id int) (err error) {
	return
}
