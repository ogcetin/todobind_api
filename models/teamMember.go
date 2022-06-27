package models

type TeamMember struct {
	TeamMemberID int    `json:"team_member_id"`
	BusinessID   int    `json:"business_id"`
	UserID       int    `json:"user_id"`
	Auth         string `json:"auth"`
}
