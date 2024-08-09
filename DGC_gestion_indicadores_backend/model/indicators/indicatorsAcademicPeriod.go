package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	"time"
)

type IndicatorsAcademicPeriod struct {
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	IndicatorTypeID  uint      `gorm:"primary_key" json:"indicator_type_id"`
	AcademicPeriodID uint      `gorm:"primary_key" json:"academic_period_id"`
	ActualValue      float64   `json:"actual_value,omitempty"`
	TargetValue      float64   `json:"target_value"`

	IndicatorType  IndicatorType                 `json:"-"`
	AcademicPeriod academicPeriod.AcademicPeriod `json:"-"`
}

func (iep IndicatorsAcademicPeriod) TableName() string {
	return model.IndicatorsSchema + ".indicators_academic_periods"
}
