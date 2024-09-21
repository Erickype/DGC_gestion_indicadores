package model

import "time"

type SpecificField struct {
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ID           uint      `gorm:"primary_key"`
	Name         string    `gorm:"size:200;not null;unique" json:"name"`
	Abbreviation string    `gorm:"size:50;" json:"abbreviation"`
	Description  string    `gorm:"size:200;" json:"description,omitempty"`
	WideFieldID  uint      `json:"wide_field_id"`

	WideField WideField `json:"-"`
}
