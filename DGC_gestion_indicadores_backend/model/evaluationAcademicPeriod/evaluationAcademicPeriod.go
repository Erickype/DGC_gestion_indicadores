package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	evaluationPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/evaluationPeriod"
	"time"
)

type EvaluationAcademicPeriod struct {
	CreatedAt          time.Time                         `json:"created_at"`
	UpdatedAt          time.Time                         `json:"updated_at"`
	EvaluationPeriodID uint                              `gorm:"primary_key;not null" json:"evaluation_period_id"`
	AcademicPeriodID   uint                              `gorm:"primary_key;not null" json:"academic_period_id"`
	EvaluationPeriod   evaluationPeriod.EvaluationPeriod `json:"-"`
	AcademicPeriod     academicPeriod.AcademicPeriod     `json:"-"`
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
