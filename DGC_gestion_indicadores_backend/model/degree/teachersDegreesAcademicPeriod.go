package model

import (
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	"time"
)

type TeachersDegreesAcademicPeriod struct {
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	AcademicPeriodID uint      `gorm:"primary_key" json:"academic_period_id"`
	TeachersDegreeID uint      `gorm:"primary_key" json:"teachers_degree_id"`

	AcademicPeriod academicPeriod.AcademicPeriod `json:"-"`
	TeachersDegree TeachersDegree                `json:"-"`
}
