package activity

import (
	activityDomain "skyshi_gethired/domain/activity"

	"gorm.io/gorm"
)

// Repository is a struct that contains the database implementation for medicine entity
type Repository struct {
	DB *gorm.DB
}

// GetAll Fetch all activity data
func (r *Repository) GetAll() (activities []activityDomain.Activity, err error) {
	resp := r.DB.Find(&activities)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return activities, nil
}

// GetByActivity Fetch all activity data
func (r *Repository) GetByActivity(activityID string) (activities []activityDomain.Activity, err error) {
	resp := r.DB.Where("activity_group_id = ?", activityID).Find(&activities)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return activities, nil
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
func (r *Repository) GetByID(id string) (activity *activityDomain.Activity, err error) {
	err = r.DB.Where("id = ?", id).First(&activity).Error
	if err != nil {
		return nil, err
	}

	return activity, nil
}

// GetActivity ... Fetch only one activity by Id
func (r *Repository) GetActivity(id string) (activity *activityDomain.Activity, err error) {
	err = r.DB.Where("id = ?", id).First(&activity).Error
	if err != nil {
		return nil, err
	}

	return activity, nil
}

// Update ... Update activity
func (r *Repository) Update(id string, activity *activityDomain.Activity) (updatedTodo *activityDomain.Activity, err error) {
	err = r.DB.Model(updatedTodo).Where("id = ?", id).Updates(activity).Error
	if err != nil {
		return nil, err
	}

	return updatedTodo, nil
}

// Delete ... Delete activity
func (r *Repository) Delete(id string) (err error) {
	err = r.DB.Delete(&activityDomain.Activity{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
