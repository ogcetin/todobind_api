package models

type TaskSchema struct {
	TaskID          int    `json:"task_id"`
	BusinessID      int    `json:"business_id"`
	SectionSchemaID int    `json:"section_schema_id"`
	Type            string `json:"type"`
	AttendantID     int    `json:"attendant_id"`
	CreationDate    string `json:"creation_date"`
	DueDate         string `json:"due_date"`
	Detail          string `json:"detail"`
}
