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
func (r *Repository) GetByID(id uint) (todo *todoDomain.Todo, err error) {
	err = r.DB.Where("id = ?", id).First(&todo).Error
	if err != nil {
		return nil, err
	}

	return todo, nil
}

// GetActivity ... Fetch only one activity by Id
func (r *Repository) GetActivity(id uint) (activity *activityDomain.Activity, err error) {
	err = r.DB.Where("id = ?", id).First(&activity).Error
	if err != nil {
		return nil, err
	}

	return activity, nil
}

// Update ... Update todo
func (r *Repository) Update(id uint, todo *todoDomain.Todo) (updatedTodo *todoDomain.Todo, err error) {
	err = r.DB.Model(updatedTodo).Where("id = ?", id).Updates(todo).Error
	if err != nil {
		return nil, err
	}

	return updatedTodo, nil
}

// Delete ... Delete todo
func (r *Repository) Delete(id uint) (err error) {
	return
}
