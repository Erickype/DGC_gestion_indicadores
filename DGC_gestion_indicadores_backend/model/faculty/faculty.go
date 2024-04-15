package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"gorm.io/gorm"
)

type Faculty struct {
	gorm.Model
	ID           uint   `gorm:"primary_key"`
	Name         string `gorm:"size:200;not null;unique" json:"name"`
	Description  string `gorm:"size:50;" json:"description"`
	Abbreviation string `gorm:"size:50;not null;unique" json:"abbreviation"`
}

func CreateFaculty(faculty *Faculty) (err error) {
	err = database.DB.Create(faculty).Error
	if err != nil {
		return err
	}
	return nil
}

func GetFaculties(faculties *[]Faculty) (err error) {
	err = database.DB.Find(faculties).Error
	if err != nil {
		return err
	}
	return nil
}

func GetFaculty(faculty *Faculty, id int) (err error) {
	err = database.DB.Where("id = ?", id).First(faculty).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateFaculty(faculty *Faculty) (err error) {
	err = database.DB.Save(faculty).Error
	if err != nil {
		return err
	}
	return nil
}
