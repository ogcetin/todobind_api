package models

type Team struct {
	TeamID     int    `json:"team_id"`
	BusinessID int    `json:"business_id"`
	CreatorID  int    `json:"creator_id"`
	Name       string `json:"name"`
}
