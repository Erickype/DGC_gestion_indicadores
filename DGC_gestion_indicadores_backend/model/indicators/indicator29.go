package model

import (
	"errors"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
)

func CalculateIndicator29(academicPeriodID int, indicator *IndicatorsAcademicPeriod) (err error) {
	var careersCount int64
	err = database.DB.Table("indicators_information.social_project_lists spl").
		Where("spl.academic_period_id = ?", academicPeriodID).
		Distinct("spl.career_id").
		Count(&careersCount).Error
	if err != nil {
		return err
	}
	if careersCount == 0 {
		return errors.New("sin carreras en lista")
	}
	var socialProjectListsCount int64
	err = database.DB.Table("indicators_information.social_project_lists spl").
		Where("spl.academic_period_id = ?", academicPeriodID).
		Count(&socialProjectListsCount).Error
	if err != nil {
		return err
	}

	indicator.ActualValue = float64(socialProjectListsCount) / float64(careersCount)
	indicator.TargetValue = model.Indicator29TargetValue
	err = RefreshIndicatorAcademicPeriod(indicator)
	if err != nil {
		return err
	}
	return nil
}
