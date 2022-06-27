package models

type Project struct {
	ProjectID  int    `json:"project_id"`
	BusinessID int    `json:"business_id"`
	TeamID     int    `json:"team_id"`
	Name       string `json:"name"`
}
