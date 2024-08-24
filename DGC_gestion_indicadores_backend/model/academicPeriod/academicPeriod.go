package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AcademicPeriod struct {
	gorm.Model
	ID           uint           `gorm:"primary_key"`
	Name         string         `gorm:"size:50;not null" json:"name"`
	Description  string         `gorm:"size:255;not null" json:"description"`
	Abbreviation string         `gorm:"size:50;unique" json:"abbreviation"`
	StartDate    datatypes.Date `gorm:"not null" json:"start_date"`
	EndDate      datatypes.Date `gorm:"not null" json:"end_date"`
}

func CreateAcademicPeriod(period *AcademicPeriod) (err error) {
	err = database.DB.Create(period).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAcademicPeriods(periods *[]AcademicPeriod) (err error) {
	err = database.DB.Order("start_date desc").Find(periods).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAcademicPeriod(period *AcademicPeriod, id int) (err error) {
	err = database.DB.Where("id = ?", id).First(period).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateAcademicPeriod(period *AcademicPeriod) (err error) {
	err = database.DB.Save(period).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteAcademicPeriod(id int) (err error) {
	err = database.DB.Delete(&AcademicPeriod{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
