package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"gorm.io/gorm"
)

type ScaledGrade struct {
	gorm.Model
	ID           uint   `gorm:"primary_key"`
	Name         string `gorm:"size:200;not null;unique" json:"name"`
	Abbreviation string `gorm:"size:50;not null;unique" json:"abbreviation"`
	Description  string `gorm:"size:200;" json:"description"`
}

func CreateScaledGrade(ScaledGrade *ScaledGrade) (err error) {
	err = database.DB.Create(ScaledGrade).Error
	if err != nil {
		return err
	}
	return nil
}

func GetScaledGrades(ScaledGrades *[]ScaledGrade) (err error) {
	err = database.DB.Find(ScaledGrades).Error
	if err != nil {
		return err
	}
	return nil
}

func GetScaledGrade(ScaledGrade *ScaledGrade, id int) (err error) {
	err = database.DB.Where("id = ?", id).First(ScaledGrade).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateScaledGrade(ScaledGrade *ScaledGrade) (err error) {
	err = database.DB.Save(ScaledGrade).Error
	if err != nil {
		return err
	}
	return nil
}
