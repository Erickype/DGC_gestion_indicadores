package model

import "time"

type CompensationFactor struct {
	ID           uint      `gorm:"primary_key"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `gorm:"size:200;not null;unique" json:"name,omitempty"`
	Abbreviation string    `gorm:"size:50;not null;unique" json:"abbreviation,omitempty"`
	Description  string    `gorm:"size:200;not null;unique" json:"description,omitempty"`
	Weight       float32   `gorm:"not null;" json:"weight"`
}
