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

func GetImpactCoefficients(impactCoefficients *[]ImpactCoefficient) (err error) {
	err = database.DB.Order("updated_at desc").Find(impactCoefficients).Error
	if err != nil {
		return err
	}
	return nil
}
