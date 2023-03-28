package activity

import "time"

// Activity is a struct that contains the data for a activity
type Activity struct {
	ID        int64  `json:"id"`
	Title     string `json:"title" example:"title activity"`
	Email     string `json:"email" example:"activity@email.com"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewActivity is a struct that contains the data for a new activity
type NewActivity struct {
	Title *string `json:"title" example:"title activity"`
	Email *string `json:"email" example:"activity@email.com"`
}

// UpdateActivity is a struct that contains the data for a update activity
type UpdateActivity struct {
	Title *string `json:"title" example:"title activity"`
}
