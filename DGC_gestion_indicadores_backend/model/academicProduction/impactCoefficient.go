package model

import "time"

type ImpactCoefficient struct {
	ID                   uint      `gorm:"primary_key"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	AcademicDatabaseID   uint      `json:"academic_database_id"`
	CompensationFactorID uint      `json:"compensation_factor_id"`

	AcademicDatabase   AcademicDatabase   `json:"-"`
	CompensationFactor CompensationFactor `json:"-"`
}
