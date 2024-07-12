package models

import "github.com/google/uuid"

type File struct {
	BaseModel
	FileID uuid.UUID `json:"file_id" db:"file_id" validate:"omitempty,uuid"`

	Name        string `json:"name"`
	FilePath    string `json:"file_path" db:"file_path" validate:"omitempty,lte=256"`
	PublicPath  string `json:"public_path" db:"public_path" validate:"omitempty,lte=256,url"`
	Size        int32  `json:"size"`
	ContentType string `json:"content_type"`
}
