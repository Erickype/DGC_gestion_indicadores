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

type PatchTeachersDegreeByTeachersDegreeIDRequest struct {
	DegreeLevelID uint   `json:"degree_level_id"`
	Name          string `json:"name"`
}

func GetTeachersDegreeByID(teachersDegree *TeachersDegree, id int) (err error) {
	err = database.DB.Where("id = ?", id).First(teachersDegree).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("título no encontrada")
		}
		return err
	}
	return nil
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

func PatchTeachersDegreeByTeachersDegreeID(request *PatchTeachersDegreeByTeachersDegreeIDRequest, id int) (err error) {
	err = database.DB.Model(&TeachersDegree{}).Where("id = ?", id).Updates(request).Error
	if err != nil {
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return errors.New("no existen las relaciones")
		}
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("título ya existe")
		}
		return err
	}
	return nil
}
