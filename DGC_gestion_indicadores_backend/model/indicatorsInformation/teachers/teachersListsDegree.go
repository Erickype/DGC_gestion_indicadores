package teachers

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	degree "github.com/Erickype/DGC_gestion_indicadores_backend/model/degree"
	"time"
)

type TeachersListsDegree struct {
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	AcademicPeriodID uint      `gorm:"primary_key" json:"academic_period_id"`
	TeachersDegreeID uint      `gorm:"primary_key" json:"teachers_degree_id"`

	AcademicPeriod academicPeriod.AcademicPeriod `json:"-"`
	TeachersDegree degree.TeachersDegree         `json:"-"`
}

func (tl *TeachersListsDegree) TableName() string {
	return model.IndicatorsInformationSchema + ".teachers_lists_degrees"
}
