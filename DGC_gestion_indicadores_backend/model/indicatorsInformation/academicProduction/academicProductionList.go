package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	academicProduction "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicProduction"
	knowledgeField "github.com/Erickype/DGC_gestion_indicadores_backend/model/knowledgeField"
	"gorm.io/datatypes"
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
	InterculturalComponent bool           `json:"intercultural_component"`

	DetailedField     knowledgeField.DetailedField         `json:"-"`
	AcademicPeriod    academicPeriod.AcademicPeriod        `json:"-"`
	ScienceMagazine   academicProduction.ScienceMagazine   `json:"-"`
	ImpactCoefficient academicProduction.ImpactCoefficient `json:"-"`
}

func (apl *AcademicProductionList) TableName() string {
	return model.IndicatorsInformationSchema + ".academic_production_lists"
}
