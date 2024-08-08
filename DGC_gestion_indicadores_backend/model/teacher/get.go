package model

import "time"

type GetByAcademicPeriodResponse struct {
	UpdatedAt        time.Time `gorm:"not null;" json:"updated_at"`
	ID               uint      `gorm:"primary_key;"`
	AcademicPeriodID uint      `gorm:"primary_key;" json:"academic_period_id"`
	PersonID         uint      `gorm:"primary_key;" json:"person_id"`
	Person           string    `gorm:"primary_key;" json:"person"`
	CareerID         uint      `gorm:"not null;" json:"career_id"`
	Career           string    `gorm:"not null;" json:"career"`
	DedicationID     uint      `gorm:"not null;" json:"dedication_id"`
	Dedication       string    `gorm:"not null;" json:"dedication"`
	ScaledGradeID    uint      `gorm:"not null;" json:"scaled_grade_id"`
	ScaledGrade      string    `gorm:"not null;" json:"scaled_grade"`
	ContractTypeID   uint      `gorm:"not null;" json:"contract_type_id"`
	ContractType     string    `gorm:"not null;" json:"contract_type"`
}
