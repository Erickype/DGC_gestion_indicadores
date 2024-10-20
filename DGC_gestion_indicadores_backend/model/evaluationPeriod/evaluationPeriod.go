package model

import (
	"errors"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type EvaluationPeriod struct {
	gorm.Model
	ID           uint           `gorm:"primary_key"`
	Name         string         `gorm:"size:50;not null" json:"name"`
	Abbreviation string         `gorm:"size:50;unique" json:"abbreviation"`
	Description  string         `gorm:"size:255;not null" json:"description"`
	StartYear    datatypes.Date `gorm:"not null" json:"start_year"`
	EndYear      datatypes.Date `gorm:"not null" json:"end_year"`
}

func CreateEvaluationPeriod(period *EvaluationPeriod) (err error) {
	err = database.DB.Create(period).Error
	if err != nil {
		return err
	}
	return nil
}

func GetEvaluationPeriods(periods *[]EvaluationPeriod) (err error) {
	err = database.DB.Order("start_year desc").Find(periods).Error
	if err != nil {
		return err
	}
	return nil
}

func GetEvaluationPeriodsByDateRange(periods *[]EvaluationPeriod, startDate datatypes.Date, endDate datatypes.Date) (err error) {
	var startDatePeriodsCount int64
	err = database.DB.Model(&EvaluationPeriod{}).
		Where("? >= start_year and ? <= end_year", startDate, startDate).
		Count(&startDatePeriodsCount).Error
	if err != nil {
		return err
	}
	if startDatePeriodsCount == 0 {
		return errors.New("sin periodos en la fecha de inicio")
	}

	var endDatePeriodsCount int64
	err = database.DB.Model(&EvaluationPeriod{}).
		Where("? >= start_year and ? <= end_year", endDate, endDate).
		Count(&endDatePeriodsCount).Error
	if err != nil {
		return err
	}
	if endDatePeriodsCount == 0 {
		return errors.New("sin periodos en la fecha de fin")
	}

	err = database.DB.Raw("? UNION DISTINCT ?",
		database.DB.Model(&EvaluationPeriod{}).
			Where("? >= start_year and ? <= end_year", startDate, startDate),
		database.DB.Model(&EvaluationPeriod{}).
			Where("? >= start_year and ? <= end_year", endDate, endDate),
	).Scan(&periods).Error

	if err != nil {
		return err
	}
	return nil
}

func GetEvaluationPeriod(period *EvaluationPeriod, id int) (err error) {
	err = database.DB.Where("id = ?", id).First(period).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateEvaluationPeriod(period *EvaluationPeriod) (err error) {
	err = database.DB.Save(period).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteEvaluationPeriod(id int) (err error) {
	err = database.DB.Delete(&EvaluationPeriod{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
