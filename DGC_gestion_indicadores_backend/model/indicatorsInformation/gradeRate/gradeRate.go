package model

import (
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

func (grl GradeRateList) TableName() string {
	return model.IndicatorsInformationSchema + ".grade_rate_lists"
}
