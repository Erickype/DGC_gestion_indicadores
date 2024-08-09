package model

import "time"

type PublicationType struct {
	ID           uint      `gorm:"primary_key;"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `gorm:"size:200;not null;unique" json:"name"`
	Abbreviation string    `gorm:"size:50;not null;unique" json:"abbreviation"`
	Description  string    `gorm:"size:200;" json:"description"`
}
