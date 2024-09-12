package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	indicatorsInformation "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/teachers"
)

func CalculateIndicator16(academicPeriodID int, indicator *IndicatorsAcademicPeriod) (err error) {
	var phdCount int64
	err = database.DB.Table("indicators_information.teachers_lists_degrees tld").
		Joins("join teachers_degrees td on tld.teachers_degree_id = td.id").
		Joins("join degree_levels dl on td.degree_level_id = dl.id").
		Where("tld.academic_period_id = ? and dl.name = ?", academicPeriodID, model.DegreeLevelPhD).
		Count(&phdCount).
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
	indicator.ActualValue = float64(phdCount) / float64(teachersCount) * 100
	indicator.TargetValue = model.Indicator16TargetValue
	err = RefreshIndicator16(indicator)
	if err != nil {
		return err
	}
	return nil
}

func RefreshIndicator16(indicator *IndicatorsAcademicPeriod) (err error) {
	err = database.DB.Save(&indicator).Error
	if err != nil {
		return err
	}
	return nil
}