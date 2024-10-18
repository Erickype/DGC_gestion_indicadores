package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	"sync"
	"time"
)

var academicPeriodIndicators = []int{16, 17, 19, 21, 29}

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

func CalculateIndicatorsByAcademicPeriod(academicPeriodID int, response *[]IndicatorAcademicPeriodJoined) (errs []error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	done := make(chan IndicatorAcademicPeriodJoined, len(academicPeriodIndicators))
	errCh := make(chan error, len(academicPeriodIndicators)) // Channel for errors

	for i := 0; i < len(academicPeriodIndicators); i++ {
		wg.Add(1)
		go func(indicatorID int) {
			defer wg.Done()

			var indicator IndicatorsAcademicPeriod
			err := CalculateIndicatorByTypeIDAndAcademicPeriod(academicPeriodID, indicatorID, &indicator)
			if err != nil {
				errCh <- err // Send error to error channel
				return
			}

			var indicatorJoined IndicatorAcademicPeriodJoined
			err = GetIndicatorByTypeIDAndAcademicPeriodJoined(academicPeriodID, indicatorID, &indicatorJoined)
			if err != nil {
				errCh <- err // Send error to error channel
				return
			}

			// Safely append to the response slice
			mu.Lock()
			*response = append(*response, indicatorJoined)
			mu.Unlock()

			// Send result to done channel
			done <- indicatorJoined
		}(academicPeriodIndicators[i])
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(done)
	close(errCh)

	// Collect all errors from the error channel
	for e := range errCh {
		if e != nil {
			errs = append(errs, e)
		}
	}

	// Return the collected errors (if any)
	return errs
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
		Order("iap.indicator_type_id").
		Scan(response).Error
	if err != nil {
		return err
	}
	return nil
}

func GetIndicatorByTypeIDAndAcademicPeriodJoined(academicPeriodID, indicatorTypeID int, response *IndicatorAcademicPeriodJoined) (err error) {
	err = database.DB.Table("indicators.indicators_academic_periods iap").
		Select(`iap.indicator_type_id,
			it.name as indicator_type,
			iap.academic_period_id,
			ap.name as academic_period,
			iap.actual_value,
			iap.target_value`).
		Joins("join indicators.indicator_types it on iap.indicator_type_id = it.id").
		Joins("join academic_periods ap on iap.academic_period_id = ap.id").
		Where("academic_period_id = ? and iap.indicator_type_id = ?", academicPeriodID, indicatorTypeID).
		Order("iap.indicator_type_id").
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
