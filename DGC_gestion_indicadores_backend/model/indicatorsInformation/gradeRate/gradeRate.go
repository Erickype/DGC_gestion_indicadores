package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	"time"
)

type GradeRateList struct {
	CreatedAt                         time.Time `json:"created_at"`
	UpdatedAt                         time.Time `json:"updated_at"`
	AcademicPeriodID                  uint      `gorm:"primary_key" json:"academic_period_id"`
	CareerID                          uint      `gorm:"primary_key" json:"career_id"`
	CountGraduatedStudents            uint      `json:"count_graduated_students"`
	CountCourtStudents                uint      `json:"count_court_students"`
	CountAdmittedMatriculatedStudents uint      `json:"count_admitted_matriculated_students"`
	CountAdmittedStudents             uint      `json:"count_admitted_students"`
}

type GradeRateListJoined struct {
	AcademicPeriodID                  uint   `json:"academic_period_id"`
	AcademicPeriod                    string `json:"academic_period"`
	CareerID                          uint   `json:"career_id"`
	Career                            string `json:"career"`
	CountGraduatedStudents            uint   `json:"count_graduated_students"`
	CountCourtStudents                uint   `json:"count_court_students"`
	CountAdmittedMatriculatedStudents uint   `json:"count_admitted_matriculated_students"`
	CountAdmittedStudents             uint   `json:"count_admitted_students"`
}

func (grl GradeRateList) TableName() string {
	return model.IndicatorsInformationSchema + ".grade_rate_lists"
}

func GetGradeRateListsByAcademicPeriod(
	academicPeriodID int, gradeRateLists *[]GradeRateListJoined) (err error) {
	err = database.DB.Table("indicators_information.grade_rate_lists grl").
		Joins("join academic_periods ap on grl.academic_period_id = ap.id").
		Joins("join careers c on grl.career_id = c.id").
		Where("where academic_period_id = ?", academicPeriodID).
		Scan(gradeRateLists).Error
	if err != nil {
		return err
	}
	return nil
}
