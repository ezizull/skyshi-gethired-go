package activity

import (
	"fmt"
	activityDomain "skyshi_gethired/domain/activity"

	"gorm.io/gorm"
)

// Repository is a struct that contains the database implementation for medicine entity
type Repository struct {
	DB *gorm.DB
}

// GetAll Fetch all activity data
func (r *Repository) GetAll() (activity []activityDomain.Activity, err error) {
	resp := r.DB.Find(&activity)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return activity, nil
}

// Create ... Insert New data
func (r *Repository) Create(newTodo *activityDomain.Activity) (*activityDomain.Activity, error) {
	resp := r.DB.Create(newTodo)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return newTodo, nil
}

// GetByID ... Fetch only one activity by Id
func (r *Repository) GetByID(id int) (activity *activityDomain.Activity, err error) {
	err = r.DB.Where("id = ?", id).First(&activity).Error
	fmt.Println("check ", activity)
	if err != nil {
		return nil, err
	}

	return activity, nil
}

// Update ... Update activity
func (r *Repository) Update(id uint, todoMap map[string]interface{}) (*activityDomain.Activity, error) {
	return nil, nil
}

// Delete ... Delete activity
func (r *Repository) Delete(id int) (err error) {
	return
}
