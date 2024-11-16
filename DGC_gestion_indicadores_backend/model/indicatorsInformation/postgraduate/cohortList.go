package model

import (
	"errors"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	"gorm.io/gorm"
	"time"
)

type CohortList struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Year      uint      `gorm:"primary_key" json:"year"`
}

func (c CohortList) TableName() string {
	return model.IndicatorsInformationSchema + ".cohort_lists"
}

func PostCohortList(cohortList *CohortList) (err error) {
	err = database.DB.Create(cohortList).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("año de cohorte ya existe")
		}
		return err
	}
	return nil
}
