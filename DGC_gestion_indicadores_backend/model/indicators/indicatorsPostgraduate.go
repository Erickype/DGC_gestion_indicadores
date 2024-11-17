package model

import (
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
