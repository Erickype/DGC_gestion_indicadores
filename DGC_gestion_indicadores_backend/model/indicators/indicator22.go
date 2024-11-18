package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
)

func CalculateIndicator22(cohortYear int, indicator *IndicatorsPostgraduate) (err error) {
	var average struct {
		avg float64
	}
	err = database.DB.Table("indicators_information.postgraduate_cohort_lists pcl").
		Select("avg(cast(pcl.graduated_students as decimal)/cast(pcl.total_students as decimal))").
		Joins("join indicators_information.postgraduate_programs pp on pcl.postgraduate_program_id = pp.id").
		Where("pcl.year = ?", cohortYear).
		Scan(&average).Error
	if err != nil {
		return
	}
	indicator.ActualValue = average.avg * 100
	indicator.TargetValue = model.Indicator22TargetValue
	err = RefreshIndicatorPostgraduate(indicator)
	if err != nil {
		return err
	}
	return nil
}
