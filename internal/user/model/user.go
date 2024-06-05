package model

import "time"

type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
	Password string    `json:"password,omitempty"`
	ApiID    int       `json:"api_id"`   // Обратите внимание на использование заглавных букв
	ApiHash  string    `json:"api_hash"` // Обратите внимание на использование заглавных букв
	Phone    string    `json:"phone"`
}

type Subscription struct {
	ID             int `json:"id"`
	UserID         int `json:"user_id"`
	SubscribedToID int `json:"subscribed_to_id"`
}
