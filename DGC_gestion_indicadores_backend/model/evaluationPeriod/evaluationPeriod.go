package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type EvaluationPeriod struct {
	gorm.Model
	ID           uint           `gorm:"primary_key"`
	Name         string         `gorm:"size:50;not null;unique" json:"name"`
	Description  string         `gorm:"size:255;not null" json:"description"`
	Abbreviation string         `gorm:"size:50;" json:"abbreviation"`
	StartYear    datatypes.Date `gorm:"not null" json:"start_year"`
	EndYear      datatypes.Date `gorm:"not null" json:"end_year"`
}

func CreateEvaluationPeriod(period *EvaluationPeriod) (err error) {
	err = database.DB.Create(period).Error
	if err != nil {
		return err
	}
	return nil
}

func GetEvaluationPeriods(periods *[]EvaluationPeriod) (err error) {
	err = database.DB.Find(periods).Error
	if err != nil {
		return err
	}
	return nil
}

func GetEvaluationPeriod(period *EvaluationPeriod, id int) (err error) {
	err = database.DB.Where("id = ?", id).First(period).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateEvaluationPeriod(period *EvaluationPeriod) (err error) {
	err = database.DB.Save(period).Error
	if err != nil {
		return err
	}
	return nil
}
