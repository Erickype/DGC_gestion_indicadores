package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"time"
)

type ScienceMagazine struct {
	ID                 uint      `gorm:"primary_key"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	Name               string    `gorm:"size:200;not null;unique" json:"name"`
	Abbreviation       string    `gorm:"size:200;not null;unique" json:"abbreviation"`
	Description        string    `gorm:"size:200" json:"description,omitempty"`
	AcademicDatabaseID uint      `json:"academic_database_id"`

	AcademicDatabase AcademicDatabase `json:"academic_database"`
}

func GetScienceMagazines(magazines *[]ScienceMagazine) (err error) {
	err = database.DB.Order("updated_at desc").Find(magazines).Error
	if err != nil {
		return err
	}
	return nil
}
