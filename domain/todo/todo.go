package todo

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID              uint      `json:"id" gorm:"primary_key"`
	ActivityGroupID uint      `json:"activity_group_id" example:"1" gorm:"index"`
	Title           string    `json:"title" example:"title todo"`
	IsActive        bool      `json:"is_active" example:"false" gorm:"default:true"`
	Priority        string    `json:"priority" example:"very-high"`
	CreatedAt       time.Time `json:"created_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoCreateTime:mili"`
	UpdatedAt       time.Time `json:"updated_at,omitempty" example:"2021-02-24 20:19:39" gorm:"autoUpdateTime:mili"`
	DeletedAt       *gorm.DeletedAt
}

func (*Todo) TableName() string {
	return "todos"
}

// Service is a interface that contains the methods for the book service
type Service interface {
	Get(int) (*Todo, error)
	GetAll() ([]*Todo, error)
	Create(*Todo) error
	GetByMap(map[string]interface{}) map[string]interface{}
	GetByID(int) (*Todo, error)
	Delete(int) error
	Update(int, map[string]interface{}) (*Todo, error)
}
