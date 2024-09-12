package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	indicatorsInformation "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/teachers"
)

func CalculateIndicator17(academicPeriodID int, indicator *IndicatorsAcademicPeriod) (err error) {
	var fullTimeCount int64
	err = database.DB.Table("indicators_information.teachers_lists tl").
		Joins("join dedications d on tl.dedication_id = d.id").
		Where("tl.academic_period_id = ? and d.name = ?", academicPeriodID, model.FullTimeDedication).
		Count(&fullTimeCount).
		Error
	if err != nil {
		return err
	}
	var teachersCount int64
	err = database.DB.Model(&indicatorsInformation.TeachersList{}).
		Where("academic_period_id = ?", academicPeriodID).
		Count(&teachersCount).Error
	if err != nil {
		return err
	}
	indicator.ActualValue = float64(fullTimeCount) / float64(teachersCount) * 100
	indicator.TargetValue = model.Indicator17TargetValue
	err = RefreshIndicator(indicator)
	if err != nil {
		return err
	}
	return nil
}
