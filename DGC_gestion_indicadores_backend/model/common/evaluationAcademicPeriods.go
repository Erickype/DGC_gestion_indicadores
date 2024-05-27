package model

import (
	"errors"

	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	acaP "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	evaAcaP "github.com/Erickype/DGC_gestion_indicadores_backend/model/evaluationAcademicPeriod"
	evaP "github.com/Erickype/DGC_gestion_indicadores_backend/model/evaluationPeriod"
	"gorm.io/gorm"
)

func CreateAcademicPeriod(acaPeriod *acaP.AcademicPeriod) (err error) {
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Create(acaPeriod).Error
		if err != nil {
			return err
		}

		var evalPeriods []evaP.EvaluationPeriod
		err = evaP.GetEvaluationPeriodsByDateRange(&evalPeriods, acaPeriod.StartDate, acaPeriod.EndDate)
		if err != nil {
			return err
		}
		if len(evalPeriods) == 0 {
			return errors.New("incorrect time period")
		}
		for _, evalPeriod := range evalPeriods {
			var evaluationAcademicPeriod evaAcaP.EvaluationAcademicPeriod
			evaluationAcademicPeriod.AcademicPeriodID = acaPeriod.ID
			evaluationAcademicPeriod.EvaluationPeriodID = evalPeriod.ID

			err = tx.Create(&evaluationAcademicPeriod).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
