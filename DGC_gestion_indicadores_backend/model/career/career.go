package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	faculty "github.com/Erickype/DGC_gestion_indicadores_backend/model/faculty"
	"gorm.io/gorm"
)

type Career struct {
	gorm.Model
	ID           uint            `gorm:"primary_key"`
	FacultyID    uint            `gorm:"not null;" json:"faculty_id"`
	Name         string          `gorm:"size:200;not null;unique" json:"name"`
	Abbreviation string          `gorm:"size:50;not null;unique" json:"abbreviation"`
	Description  string          `gorm:"size:200;" json:"description"`
	Faculty      faculty.Faculty `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func CreateCareer(Career *Career) (err error) {
	err = database.DB.Create(Career).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCareers(careers *[]Career) (err error) {
	err = database.DB.Find(careers).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCareer(Career *Career, id int) (err error) {
	err = database.DB.Where("id = ?", id).First(Career).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateCareer(Career *Career) (err error) {
	err = database.DB.Save(Career).Error
	if err != nil {
		return err
	}
	return nil
}
