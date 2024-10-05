package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	evaluationAcademicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/evaluationAcademicPeriod"
)

func CalculateIndicator26(evaluationPeriodID int, indicator *IndicatorsEvaluationPeriod) (err error) {
	var academicPeriodIds []int
	err = getAcademicPeriodsInEvaluationPeriod(evaluationPeriodID, &academicPeriodIds)
	if err != nil {
		return err
	}

	var academicPublication float64
	err = calculateIndicator26AcademicPublication(academicPeriodIds, &academicPublication)
	if err != nil {
		return err
	}

	indicator.ActualValue = academicPublication
	indicator.TargetValue = model.Indicator26TargetValue
	err = RefreshIndicatorEvaluationPeriod(indicator)
	if err != nil {
		return err
	}
	return nil
}

func calculateIndicator26AcademicPublication(academicPeriodIDs []int, academicPublication *float64) (err error) {
	var weights []struct {
		Weight                 float64 `json:"weight"`
		InterculturalComponent bool    `json:"intercultural_component"`
	}
	err = database.DB.Table("indicators_information.academic_production_lists apl").
		Joins("join impact_coefficients ic on apl.impact_coefficient_id = ic.id").
		Joins("join compensation_factors cf on ic.compensation_factor_id = cf.id").
		Where("apl.academic_period_id in ?", academicPeriodIDs).
		Scan(&weights).Error

	*academicPublication = 0
	for _, weight := range weights {
		var weightedValue float64 = 0
		if weight.InterculturalComponent {
			weightedValue = weight.Weight + 0.21
			if weightedValue > 1 {
				weightedValue = 1
			}
		} else {
			weightedValue = weight.Weight
		}
		*academicPublication += weightedValue
	}
	return nil
}

func getAcademicPeriodsInEvaluationPeriod(evaluationPeriodID int, academicPeriodIDs *[]int) (err error) {
	err = database.DB.Model(&evaluationAcademicPeriod.EvaluationAcademicPeriod{}).
		Select("academic_period_id").
		Where("evaluation_period_id = ?", evaluationPeriodID).
		Scan(&academicPeriodIDs).Error
	if err != nil {
		return err
	}
	return nil
}
