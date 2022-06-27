package models

type Task struct {
	TaskID       int    `json:"task_id"`
	ProjectID    int    `json:"project_id"`
	BusinessID   int    `json:"business_id"`
	CreatorID    int    `json:"creator_id"`
	AttendantID  int    `json:"attendant_id"`
	CreationDate string `json:"creation_date"`
	DueDate      string `json:"due_date"`
	Detail       string `json:"detail"`
}
