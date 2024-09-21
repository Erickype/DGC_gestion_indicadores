package model

import "time"

type DetailedField struct {
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	ID              uint      `gorm:"primary_key"`
	Name            string    `gorm:"size:200;not null;unique" json:"name"`
	Abbreviation    string    `gorm:"size:200" json:"abbreviation"`
	Description     string    `gorm:"size:200;" json:"description,omitempty"`
	SpecificFieldID uint      `json:"specific_field_id"`

	SpecificField SpecificField `json:"-"`
}
