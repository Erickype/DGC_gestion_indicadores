package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	"time"
)

type IndicatorsAcademicPeriod struct {
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	IndicatorTypeID  uint      `gorm:"primary_key" json:"indicator_type_id"`
	AcademicPeriodID uint      `gorm:"primary_key" json:"academic_period_id"`
	ActualValue      float64   `json:"actual_value"`
	TargetValue      float64   `json:"target_value"`

	IndicatorType  IndicatorType                 `json:"-"`
	AcademicPeriod academicPeriod.AcademicPeriod `json:"-"`
}

type IndicatorAcademicPeriodJoined struct {
	IndicatorTypeID  uint     `json:"indicator_type_id"`
	IndicatorType    string   `json:"indicator_type"`
	AcademicPeriodID uint     `json:"academic_period_id"`
	AcademicPeriod   string   `json:"academic_period"`
	ActualValue      *float64 `json:"actual_value,omitempty"`
	TargetValue      float64  `json:"target_value"`
}

func (iep IndicatorsAcademicPeriod) TableName() string {
	return model.IndicatorsSchema + ".indicators_academic_periods"
}

func CalculateIndicatorByTypeIDAndAcademicPeriod(academicPeriodID, indicatorTypeID int, response *IndicatorsAcademicPeriod) error {
	response.IndicatorTypeID = uint(indicatorTypeID)
	response.AcademicPeriodID = uint(academicPeriodID)
	switch indicatorTypeID {
	case model.Indicator16:
		err := CalculateIndicator16(academicPeriodID, response)
		if err != nil {
			return err
		}
	case model.Indicator17:
		err := CalculateIndicator17(academicPeriodID, response)
		if err != nil {
			return err
		}
	case model.Indicator19:
		err := CalculateIndicator19(academicPeriodID, response)
		if err != nil {
			return err
		}
	case model.Indicator21:
		err := CalculateIndicator21(academicPeriodID, response)
		if err != nil {
			return err
		}
	case model.Indicator29:
		err := CalculateIndicator29(academicPeriodID, response)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetIndicatorsByAcademicPeriod(academicPeriodID int, response *[]IndicatorAcademicPeriodJoined) (err error) {
	err = database.DB.Table("indicators.indicators_academic_periods iap").
		Select(`iap.indicator_type_id,
			it.name as indicator_type,
			iap.academic_period_id,
			ap.name as academic_period,
			iap.actual_value,
			iap.target_value`).
		Joins("join indicators.indicator_types it on iap.indicator_type_id = it.id").
		Joins("join academic_periods ap on iap.academic_period_id = ap.id").
		Where("academic_period_id = ?", academicPeriodID).
		Scan(response).Error
	if err != nil {
		return err
	}
	return nil
}

func RefreshIndicatorAcademicPeriod(indicator *IndicatorsAcademicPeriod) (err error) {
	err = database.DB.Save(&indicator).Error
	if err != nil {
		return err
	}
	return nil
}
