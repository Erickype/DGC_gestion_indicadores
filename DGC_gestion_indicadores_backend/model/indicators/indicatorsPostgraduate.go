package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	postgraduate "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/postgraduate"
	"time"
)

type IndicatorsPostgraduate struct {
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	IndicatorTypeID uint      `gorm:"primary_key" json:"indicator_type_id"`
	CohortListYear  uint      `gorm:"primary_key" json:"year"`
	ActualValue     float64   `json:"actual_value"`
	TargetValue     float64   `json:"target_value"`

	IndicatorType IndicatorType           `json:"-"`
	CohortList    postgraduate.CohortList `json:"-"`
}

type IndicatorsPostgraduateJoined struct {
	IndicatorTypeID uint     `json:"indicator_type_id"`
	IndicatorType   string   `json:"indicator_type"`
	CohortListYear  uint     `json:"cohort_list_year"`
	ActualValue     *float64 `json:"actual_value,omitempty"`
	TargetValue     float64  `json:"target_value"`
}

func (iep IndicatorsPostgraduate) TableName() string {
	return model.IndicatorsSchema + ".indicators_postgraduates"
}

func RefreshIndicatorPostgraduate(indicator *IndicatorsPostgraduate) (err error) {
	err = database.DB.Save(indicator).Error
	return nil
}

func CalculateIndicatorByTypeIDAndCohortYear(indicatorTypeID, cohortYear int, response *IndicatorsPostgraduate) (err error) {
	switch indicatorTypeID {
	case model.Indicator22:
		err = CalculateIndicator22(cohortYear, response)
		if err != nil {
			return err
		}
		break
	}
	return nil
}

func GetIndicatorsByPostgraduateCohortYear(
	cohortYear int, response *[]IndicatorsPostgraduateJoined) (err error) {
	err = database.DB.Table("indicators.indicators_postgraduates ip").
		Select(`ip.indicator_type_id,
				it.name as indicator_type,
				ip.cohort_list_year,
				ip.actual_value,
				ip.target_value`).
		Joins("join indicators.indicator_types it on ip.indicator_type_id = it.id").
		Where("ip.cohort_list_year = ?", cohortYear).
		Order("ip.indicator_type_id").
		Scan(response).Error
	if err != nil {
		return err
	}
	return nil
}
