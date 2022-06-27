package models

type SectionSchema struct {
	SectionSchemaID int    `json:"section_schema_id"`
	BusinessID      int    `json:"business_id"`
	Type            string `json:"type"`
	Name            string `json:"name"`
}
