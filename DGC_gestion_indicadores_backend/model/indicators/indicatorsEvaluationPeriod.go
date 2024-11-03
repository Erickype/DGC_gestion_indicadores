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

type IndicatorEvaluationPeriodJoined struct {
	IndicatorTypeID    uint     `json:"indicator_type_id"`
	IndicatorType      string   `json:"indicator_type"`
	EvaluationPeriodID uint     `json:"evaluation_period_id"`
	EvaluationPeriod   string   `json:"evaluation_period"`
	ActualValue        *float64 `json:"actual_value,omitempty"`
	TargetValue        float64  `json:"target_value"`
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
		break
	case model.Indicator25:
		err := CalculateIndicator25(evaluationPeriodID, response)
		if err != nil {
			return err
		}
		break
	}
	return nil
}

func GetIndicatorByTypeIDAndEvaluationPeriodJoined(
	evaluationPeriodID, indicatorTypeID int, indicatorEvaluationPeriodJoined *IndicatorEvaluationPeriodJoined) (err error) {
	err = database.DB.Table("indicators.indicators_evaluation_periods iep").
		Select(`iep.indicator_type_id,
				it.name as indicator_type,
				iep.evaluation_period_id,
				ep.name as evaluation_period,
				iep.actual_value,
				iep.target_value`).
		Joins("join indicators.indicator_types it on iep.indicator_type_id = it.id").
		Joins("join evaluation_periods ep on iep.evaluation_period_id = ep.id").
		Where("iep.evaluation_period_id = ?", evaluationPeriodID).
		Where("iep.indicator_type_id = ?", indicatorTypeID).
		Scan(&indicatorEvaluationPeriodJoined).Error
	if err != nil {
		return err
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
