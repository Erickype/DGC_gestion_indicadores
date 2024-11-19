package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	postgraduate "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/postgraduate"
	"sync"
	"time"
)

var postgraduateIndicators = []int{22}

type IndicatorsPostgraduate struct {
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	IndicatorTypeID uint      `gorm:"primary_key" json:"indicator_type_id"`
	CohortListYear  uint      `gorm:"primary_key" json:"cohort_list_year"`
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

func GetIndicatorByTypeIDAndPostgraduateCohortYearJoined(
	cohortYear, indicatorTypeID int, response *IndicatorsPostgraduateJoined) (err error) {
	err = database.DB.Table("indicators.indicators_postgraduates ip").
		Select(`ip.indicator_type_id,
				it.name as indicator_type,
				ip.cohort_list_year,
				ip.actual_value,
				ip.target_value`).
		Joins("join indicators.indicator_types it on ip.indicator_type_id = it.id").
		Where("ip.cohort_list_year = ?", cohortYear).
		Where("ip.indicator_type_id = ?", indicatorTypeID).
		Scan(response).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCalculateByPostgraduateCohortYear(cohortYear int, response *[]IndicatorsPostgraduateJoined) (errs []error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	done := make(chan IndicatorsPostgraduateJoined, len(postgraduateIndicators))
	errCh := make(chan error, len(postgraduateIndicators)) // Channel for errors

	for i := 0; i < len(postgraduateIndicators); i++ {
		wg.Add(1)
		go func(indicatorID int) {
			defer wg.Done()

			var indicator IndicatorsPostgraduate
			err := CalculateIndicatorByTypeIDAndCohortYear(cohortYear, indicatorID, &indicator)
			if err != nil {
				errCh <- err // Send error to error channel
				return
			}

			var indicatorJoined IndicatorsPostgraduateJoined
			err = GetIndicatorByTypeIDAndPostgraduateCohortYearJoined(cohortYear, indicatorID, &indicatorJoined)
			if err != nil {
				errCh <- err // Send error to error channel
				return
			}

			// Safely append to the response slice
			mu.Lock()
			*response = append(*response, indicatorJoined)
			mu.Unlock()

			done <- indicatorJoined
		}(evaluationPeriodIndicators[i])
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
