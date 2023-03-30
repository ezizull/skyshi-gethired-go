package todo

import (
	activityDomain "skyshi_gethired/domain/activity"
	todoDomain "skyshi_gethired/domain/todo"

	"gorm.io/gorm"
)

// Repository is a struct that contains the database implementation for medicine entity
type Repository struct {
	DB *gorm.DB
}

// GetAll Fetch all todo data
func (r *Repository) GetAll() (todos []todoDomain.Todo, err error) {
	resp := r.DB.Find(&todos)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return todos, nil
}

// GetByActivity Fetch all todo data
func (r *Repository) GetByActivity(activityID string) (todos []todoDomain.Todo, err error) {
	resp := r.DB.Where("activity_group_id = ?", activityID).Find(&todos)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return todos, nil
}

// Create ... Insert New data
func (r *Repository) Create(newTodo *todoDomain.Todo) (*todoDomain.Todo, error) {
	resp := r.DB.Create(newTodo)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return newTodo, nil
}

// GetByID ... Fetch only one todo by Id
func (r *Repository) GetByID(id string) (todo *todoDomain.Todo, err error) {
	resp := r.DB.Where("id = ?", id).First(&todo)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return todo, nil
}

// GetActivity ... Fetch only one activity by Id
func (r *Repository) GetActivity(id uint) (activity *activityDomain.Activity, err error) {
	resp := r.DB.Where("id = ?", id).First(&activity)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return activity, nil
}

// Update ... Update todo
func (r *Repository) Update(id string, todo *todoDomain.Todo) (updatedTodo *todoDomain.Todo, err error) {
	resp := r.DB.Model(updatedTodo).Where("id = ?", id).Updates(todo)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return updatedTodo, nil
}

// Delete ... Delete todo
func (r *Repository) Delete(id string) (err error) {
	resp := r.DB.Delete(&todoDomain.Todo{}, id)
	if resp.Error != nil {
		return resp.Error
	}
	return nil
}
