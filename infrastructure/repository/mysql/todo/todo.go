package todo

import (
	domainTodo "skyshi_gethired/domain/todo"

	"gorm.io/gorm"
)

// Repository is a struct that contains the database implementation for medicine entity
type Repository struct {
	DB *gorm.DB
}

// GetAll Fetch all todo data
func (r *Repository) GetAll() (todos []domainTodo.Todo, err error) {
	resp := r.DB.Find(&todos)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return todos, nil
}

// GetByActivity Fetch all todo data
func (r *Repository) GetByActivity(activityID string) (todos []domainTodo.Todo, err error) {
	resp := r.DB.Where("activity_group_id = ?", activityID).Find(&todos)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return todos, nil
}

// Create ... Insert New data
func (r *Repository) Create(newTodo *domainTodo.Todo) (*domainTodo.Todo, error) {
	resp := r.DB.Create(newTodo)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return newTodo, nil
}

// GetByID ... Fetch only one todo by Id
func (r *Repository) GetByID(id int) (todo *domainTodo.Todo, err error) {
	err = r.DB.Where("id = ?", id).First(&todo).Error
	if err != nil {
		return nil, err
	}

	return todo, nil
}

// Update ... Update todo
func (r *Repository) Update(id uint, todoMap map[string]interface{}) (*domainTodo.Todo, error) {
	return nil, nil
}

// Delete ... Delete todo
func (r *Repository) Delete(id int) (err error) {
	return
}
