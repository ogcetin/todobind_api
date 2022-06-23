package models

type User struct {
	UserID     int    `json:"user_id"`
	BusinessID int    `json:"business_id"`
	Name       string `json:"name"`
	Status     string `json:"status"`
}
