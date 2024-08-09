package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	evaluationPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/evaluationPeriod"
	"time"
)

type IndicatorsEvaluationPeriod struct {
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	IndicatorTypeID    uint      `gorm:"primary_key" json:"indicator_type_id"`
	EvaluationPeriodID uint      `gorm:"primary_key" json:"evaluation_period_id"`
	ActualValue        float64   `json:"actual_value,omitempty"`
	TargetValue        float64   `json:"target_value"`

	IndicatorType    IndicatorType                     `json:"-"`
	EvaluationPeriod evaluationPeriod.EvaluationPeriod `json:"-"`
}

func (iep IndicatorsEvaluationPeriod) TableName() string {
	return model.IndicatorsSchema + ".indicators_evaluation_periods"
}
