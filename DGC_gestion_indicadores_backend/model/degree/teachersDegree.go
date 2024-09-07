package model

import (
	"errors"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	teacher "github.com/Erickype/DGC_gestion_indicadores_backend/model/teacher"
	"gorm.io/gorm"
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

func PostTeachersDegree(degree *TeachersDegree) (err error) {
	err = database.DB.Create(degree).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("título ya registrado")
		}
		return err
	}
	return nil
}
