package model

import (
	teacher "github.com/Erickype/DGC_gestion_indicadores_backend/model/teacher"
	"time"
)

type TeachersDegree struct {
	ID            uint      `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DegreeLevelID uint      `json:"degree_level_id"`
	TeacherID     uint      `json:"teacher_id"`
	Name          string    `gorm:"size:200;not null;unique" json:"name,omitempty"`

	DegreeLevel DegreeLevel     `json:"-"`
	Teacher     teacher.Teacher `json:"-"`
}
