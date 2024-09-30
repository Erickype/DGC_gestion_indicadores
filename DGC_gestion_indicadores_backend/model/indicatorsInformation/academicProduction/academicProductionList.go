package model

import (
	"errors"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	academicProduction "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicProduction"
	knowledgeField "github.com/Erickype/DGC_gestion_indicadores_backend/model/knowledgeField"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type AcademicProductionList struct {
	CreatedAt              time.Time      `json:"created_at"`
	UpdatedAt              time.Time      `json:"updated_at"`
	ID                     uint           `gorm:"primary_key"`
	DOI                    string         `gorm:"size:1000;not null;unique" json:"DOI"`
	PublicationDate        datatypes.Date `gorm:"not null" json:"publication_date"`
	PublicationName        string         `gorm:"size:1000;not null" json:"publication_name"`
	AcademicPeriodID       uint           `json:"academic_period_id"`
	DetailedFieldID        uint           `json:"detailed_field_id"`
	ScienceMagazineID      uint           `json:"science_magazine_id"`
	ImpactCoefficientID    uint           `json:"impact_coefficient_id"`
	InterculturalComponent *bool          `json:"intercultural_component"`

	DetailedField     knowledgeField.DetailedField         `json:"-"`
	AcademicPeriod    academicPeriod.AcademicPeriod        `json:"-"`
	ScienceMagazine   academicProduction.ScienceMagazine   `json:"-"`
	ImpactCoefficient academicProduction.ImpactCoefficient `json:"-"`
}

func (apl *AcademicProductionList) TableName() string {
	return model.IndicatorsInformationSchema + ".academic_production_lists"
}

func GetAcademicProductionListByID(id int, response *AcademicProductionList) (err error) {
	err = database.DB.First(&response, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("no existe la publicación en el periodo")
		}
		return err
	}
	return nil
}

func PostAcademicProductionList(request *AcademicProductionList) (err error) {
	err = database.DB.Create(request).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("publicación académica en lista ya existe")
		}
		return err
	}
	return nil
}

func PatchAcademicProductionList(academicProductionList *AcademicProductionList) (err error) {
	err = database.DB.Model(&AcademicProductionList{}).
		Where("id = ?",
			academicProductionList.ID).
		Updates(academicProductionList).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("publicación académica en lista ya existe")
		}
		return err
	}
	return nil
}
