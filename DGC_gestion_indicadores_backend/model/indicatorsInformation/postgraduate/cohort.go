package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	"time"
)

type PostgraduateCohortList struct {
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	PostgraduateProgramID uint      `gorm:"primary_key" json:"program_id"`
	Year                  uint      `gorm:"primary_key" json:"year"`
	Name                  string    `json:"name"`
	GraduatedStudents     uint      `json:"graduated_students"`
	TotalStudents         uint      `json:"total_students"`

	PostgraduateProgram PostgraduateProgram `json:"-"`
}

func (c PostgraduateCohortList) TableName() string {
	return model.IndicatorsInformationSchema + ".postgraduate_cohort_lists"
}
