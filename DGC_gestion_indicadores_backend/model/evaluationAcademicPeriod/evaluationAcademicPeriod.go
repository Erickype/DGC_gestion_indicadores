package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	evaluationPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/evaluationPeriod"
	"gorm.io/gorm"
)

type EvaluationAcademicPeriod struct {
	gorm.Model
	EvaluationPeriodID uint `gorm:"primary_key;not null" json:"evaluation_period_id"`
	AcademicPeriodID   uint `gorm:"primary_key;not null" json:"academic_period_id"`
	EvaluationPeriod   evaluationPeriod.EvaluationPeriod
	AcademicPeriod     academicPeriod.AcademicPeriod
}

func CreateEvaluationAcademicPeriod(period *EvaluationAcademicPeriod) (err error) {
	err = database.DB.Create(period).Error
	if err != nil {
		return err
	}
	return nil
}
