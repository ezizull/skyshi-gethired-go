package activity

import "time"

// Activity is a struct that contains the todo model
type Activity struct {
	ID        uint   `json:"id"`
	Title     string `json:"title" example:"title activity"`
	Email     string `json:"email" example:"activity@email.com"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName overrides the table name used by User to `users`
func (*Activity) TableName() string {
	return "activities"
}
