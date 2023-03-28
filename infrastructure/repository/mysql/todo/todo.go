package todo

import (
	"encoding/json"
	"log"
	domainErrors "skyshi_gethired/domain/errors"
	domainTodo "skyshi_gethired/domain/todo"

	"gorm.io/gorm"
)

// Repository is a struct that contains the database implementation for medicine entity
type Repository struct {
	DB *gorm.DB
}

// GetAll Fetch all todo data
func (r *Repository) GetAll(page int64, limit int64) (todos *[]domainTodo.Todo, err error) {
	resp := r.DB.Find(todos)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return todos, nil
}

// Create ... Insert New data
func (r *Repository) Create(newTodo *domainTodo.Todo) (createdTodo *domainTodo.Todo, err error) {
	todo := fromDomainMapper(newTodo)

	tx := r.DB.Create(todo)

	if tx.Error != nil {
		byteErr, _ := json.Marshal(tx.Error)
		var newError domainErrors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return
		}
		switch newError.Number {
		case 1062:
			err = domainErrors.NewAppErrorWithType(domainErrors.ResourceAlreadyExists)
		default:
			err = domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
		}
		return
	}

	createdTodo = todo.toDomainMapper()
	return
}

// GetByID ... Fetch only one todo by Id
func (r *Repository) GetByID(id int) (*domainTodo.Todo, error) {
	var todo Todo
	err := r.DB.Where("id = ?", id).First(&todo).Error

	if err != nil {
		switch err.Error() {
		case gorm.ErrRecordNotFound.Error():
			err = domainErrors.NewAppErrorWithType(domainErrors.NotFound)
		default:
			err = domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
		}
		return &domainTodo.Todo{}, err
	}

	return todo.toDomainMapper(), nil
}

// GetOneByMap ... Fetch only one todo by Map
func (r *Repository) GetOneByMap(todoMap map[string]interface{}) (*domainTodo.Todo, error) {
	var todo Todo

	err := r.DB.Where(todoMap).Limit(1).Find(&todo).Error
	if err != nil {
		err = domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
		return nil, err
	}
	return todo.toDomainMapper(), err
}

// Update ... Update todo
func (r *Repository) Update(id uint, todoMap map[string]interface{}) (*domainTodo.Todo, error) {
	var todo Todo

	todo.ID = id
	err := r.DB.Model(&todo).
		Select("title", "author", "description").
		Updates(todoMap).Error

	// err = config.DB.Save(todo).Error
	if err != nil {
		byteErr, _ := json.Marshal(err)
		var newError domainErrors.GormErr
		err = json.Unmarshal(byteErr, &newError)
		if err != nil {
			return &domainTodo.Todo{}, err
		}
		switch newError.Number {
		case 1062:
			err = domainErrors.NewAppErrorWithType(domainErrors.ResourceAlreadyExists)
			return &domainTodo.Todo{}, err

		default:
			err = domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
			return &domainTodo.Todo{}, err
		}
	}

	err = r.DB.Where("id = ?", id).First(&todo).Error

	return todo.toDomainMapper(), err
}

// Delete ... Delete todo
func (r *Repository) Delete(id int) (err error) {
	tx := r.DB.Delete(&domainTodo.Todo{}, id)

	log.Println("check ", tx)
	if tx.Error != nil {
		err = domainErrors.NewAppErrorWithType(domainErrors.UnknownError)
		return
	}

	if tx.RowsAffected == 0 {
		err = domainErrors.NewAppErrorWithType(domainErrors.NotFound)
	}

	return
}
