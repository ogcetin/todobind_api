package models

type Business struct {
	BusinessID int    `json:"business_id"`
	Name       string `json:"name"`
	Status     string `json:"status"`
}
