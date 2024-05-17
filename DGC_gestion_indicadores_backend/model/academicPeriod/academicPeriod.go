package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	period "github.com/Erickype/DGC_gestion_indicadores_backend/model/evaluationPeriod"
	"gorm.io/gorm"
)

type AcademicPeriod struct {
	gorm.Model
	ID                 uint                    `gorm:"primary_key"`
	EvaluationPeriodID uint                    `gorm:"not null;" json:"evaluation_period_id"`
	Name               string                  `gorm:"size:50;not null;unique" json:"name"`
	Description        string                  `gorm:"size:255;not null" json:"description"`
	EvaluationPeriod   period.EvaluationPeriod `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func CreateAcademicPeriod(period *AcademicPeriod) (err error) {
	err = database.DB.Create(period).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAcademicPeriods(periods *[]AcademicPeriod) (err error) {
	err = database.DB.Find(periods).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAcademicPeriod(period *AcademicPeriod, id int) (err error) {
	err = database.DB.Where("id = ?", id).First(period).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateAcademicPeriod(period *AcademicPeriod) (err error) {
	err = database.DB.Save(period).Error
	if err != nil {
		return err
	}
	return nil
}
