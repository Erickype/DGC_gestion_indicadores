package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	evaluationPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/evaluationPeriod"
	"gorm.io/gorm"
)

type EvaluationAcademicPeriod struct {
	gorm.Model
	EvaluationPeriodID uint                              `gorm:"primary_key;not null" json:"evaluation_period_id"`
	AcademicPeriodID   uint                              `gorm:"primary_key;not null" json:"academic_period_id"`
	EvaluationPeriod   evaluationPeriod.EvaluationPeriod `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	AcademicPeriod     academicPeriod.AcademicPeriod     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func CreateEvaluationAcademicPeriod(period *EvaluationAcademicPeriod) (err error) {
	err = database.DB.Create(period).Error
	if err != nil {
		return err
	}
	return nil
}

func GetEvaluationAcademicPeriodByAcademicPeriod(evaAcaPeriods *[]EvaluationAcademicPeriod, idAcademicPeriod int) (err error) {
	err = database.DB.Where("academic_period_id = ?", idAcademicPeriod).Find(evaAcaPeriods).Error
	if err != nil {
		return err
	}
	return nil
}
