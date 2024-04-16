package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"gorm.io/gorm"
)

type Dedication struct {
	gorm.Model
	ID           uint   `gorm:"primary_key"`
	Name         string `gorm:"size:200;not null;unique" json:"name"`
	Abbreviation string `gorm:"size:50;not null;unique" json:"abbreviation"`
	Description  string `gorm:"size:200;" json:"description"`
}

func CreateDedication(Dedication *Dedication) (err error) {
	err = database.DB.Create(Dedication).Error
	if err != nil {
		return err
	}
	return nil
}

func GetDedications(Dedications *[]Dedication) (err error) {
	err = database.DB.Find(Dedications).Error
	if err != nil {
		return err
	}
	return nil
}

func GetDedication(Dedication *Dedication, id int) (err error) {
	err = database.DB.Where("id = ?", id).First(Dedication).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateDedication(Dedication *Dedication) (err error) {
	err = database.DB.Save(Dedication).Error
	if err != nil {
		return err
	}
	return nil
}
