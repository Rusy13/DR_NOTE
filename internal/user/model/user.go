package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
}

type Subscription struct {
	ID             int `json:"id"`
	UserID         int `json:"user_id"`
	SubscribedToID int `json:"subscribed_to_id"`
}
