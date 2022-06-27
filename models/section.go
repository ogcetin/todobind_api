package models

type Section struct {
	SectionID  int    `json:"section_id"`
	ProjectID  int    `json:"project_id"`
	BusinessID int    `json:"business_id"`
	Name       string `json:"name"`
}
