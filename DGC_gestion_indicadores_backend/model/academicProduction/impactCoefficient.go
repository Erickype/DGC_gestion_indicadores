package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"time"
)

type ImpactCoefficient struct {
	ID                   uint      `gorm:"primary_key"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	AcademicDatabaseID   uint      `json:"academic_database_id"`
	CompensationFactorID uint      `json:"compensation_factor_id"`

	AcademicDatabase   AcademicDatabase   `json:"-"`
	CompensationFactor CompensationFactor `json:"-"`
}

type ImpactCoefficientJoined struct {
	ID                   uint
	AcademicDatabaseID   uint   `json:"academic_database_id,omitempty"`
	AcademicDatabase     string `json:"academic_database,omitempty"`
	CompensationFactorID uint   `json:"compensation_factor_id,omitempty"`
	CompensationFactor   string `json:"compensation_factor,omitempty"`
}

func GetImpactCoefficients(impactCoefficients *[]ImpactCoefficientJoined) (err error) {
	err = database.DB.Order("ic.updated_at desc").
		Table("impact_coefficients ic").
		Select(`ic.id,
			ic.academic_database_id,
			ad.name as academic_database,
			ic.compensation_factor_id,
			cf.name as compensation_factor`).
		Joins("join academic_databases ad on ad.id = ic.academic_database_id").
		Joins("join compensation_factors cf on cf.id = ic.compensation_factor_id").
		Scan(impactCoefficients).Error
	if err != nil {
		return err
	}
	return nil
}
