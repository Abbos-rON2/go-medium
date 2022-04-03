package models

import "time"

type Subscription struct {
	ID        int       `json:"-"`
	UserID    int       `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
