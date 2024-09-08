package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"time"
)

type DegreeLevel struct {
	ID           uint      `gorm:"primary_key"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `gorm:"size:200;not null;unique" json:"name"`
	Abbreviation string    `gorm:"size:50;not null;unique" json:"abbreviation"`
	Description  string    `gorm:"size:200;" json:"description,omitempty"`
}

func GetDegreeLevels(degreeLevel *[]DegreeLevel) (err error) {
	err = database.DB.Order("created_at desc").Find(degreeLevel).Error
	if err != nil {
		return err
	}
	return nil
}
