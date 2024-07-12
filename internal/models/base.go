package models

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	UserID    uuid.UUID `json:"user_id" db:"user_id" redis:"user_id" validate:"omitempty"`
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
