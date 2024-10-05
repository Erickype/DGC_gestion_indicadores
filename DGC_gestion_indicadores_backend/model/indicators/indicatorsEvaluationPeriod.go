package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
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

func CalculateIndicatorByTypeIDAndEvaluationPeriod(evaluationPeriodID, indicatorTypeID int, response *IndicatorsEvaluationPeriod) error {
	response.IndicatorTypeID = uint(indicatorTypeID)
	response.EvaluationPeriodID = uint(evaluationPeriodID)
	switch indicatorTypeID {
	case model.Indicator26:
		err := CalculateIndicator26(evaluationPeriodID, response)
		if err != nil {
			return err
		}
	}
	return nil
}

func RefreshIndicatorEvaluationPeriod(indicator *IndicatorsEvaluationPeriod) (err error) {
	err = database.DB.Save(&indicator).Error
	if err != nil {
		return err
	}
	return nil
}
