package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	academicProduction "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicProduction"
	"gorm.io/datatypes"
	"time"
)

type AcademicProductionList struct {
	DOI                    string         `gorm:"primary_key;size:1000" json:"DOI"`
	CreatedAt              time.Time      `json:"created_at"`
	UpdatedAt              time.Time      `json:"updated_at"`
	AcademicPeriodID       uint           `json:"academic_period_id"`
	PublicationDate        datatypes.Date `gorm:"not null" json:"publication_date"`
	PublicationName        string         `gorm:"size:1000;not null" json:"publication_name"`
	PublicationTypeID      uint16         `json:"publication_type_id"`
	ScienceMagazineID      uint           `json:"science_magazine_id"`
	ImpactCoefficientID    uint           `json:"impact_coefficient_id"`
	InterculturalComponent bool           `json:"intercultural_component"`

	PublicationType   academicProduction.PublicationType   `json:"-"`
	AcademicPeriod    academicPeriod.AcademicPeriod        `json:"-"`
	ScienceMagazine   academicProduction.ScienceMagazine   `json:"-"`
	ImpactCoefficient academicProduction.ImpactCoefficient `json:"-"`
}

func (apl *AcademicProductionList) TableName() string {
	return model.IndicatorsInformationSchema + ".academic_production_lists"
}
