package model

import (
	"errors"
	"fmt"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
)

func CalculateIndicator21(academicPeriodID int, indicator *IndicatorsAcademicPeriod) (err error) {
	var facultiesAverage []struct {
		Avg float64 `json:"avg"`
	}
	err = database.DB.Table("indicators_information.grade_rate_lists grl").
		Select("avg(cast(grl.count_graduated_students as decimal) / cast(grl.count_court_students as decimal)) as avg").
		Joins("join careers c on grl.career_id = c.id").
		Joins("join faculties f on c.faculty_id = f.id").
		Where("grl.academic_period_id = ?", academicPeriodID).
		Group("f.id").
		Scan(&facultiesAverage).Error
	if err != nil {
		return
	}
	if len(facultiesAverage) >= 0 {
		return errors.New(fmt.Sprintf("indicador %d: sin valores para cálculo de tasas", indicator.IndicatorTypeID))
	}

	var sum float64
	for _, average := range facultiesAverage {
		sum += average.Avg
	}

	indicator.ActualValue = sum / float64(len(facultiesAverage)) * 100
	indicator.TargetValue = model.Indicator21TargetValue
	err = RefreshIndicatorAcademicPeriod(indicator)
	if err != nil {
		return err
	}
	return nil
}
