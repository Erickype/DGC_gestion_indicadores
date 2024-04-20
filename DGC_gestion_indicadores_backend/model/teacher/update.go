package model

type Update struct {
	ID               uint `gorm:"primary_key;"`
	AcademicPeriodID uint `gorm:"primary_key;" json:"academic_period_id"`
	PersonID         uint `gorm:"primary_key;" json:"person_id"`
	CareerID         uint `gorm:"not null;" json:"career_id"`
	DedicationID     uint `gorm:"not null;" json:"dedication_id"`
	ScaledGradeID    uint `gorm:"not null;" json:"scaled_grade_id"`
}
