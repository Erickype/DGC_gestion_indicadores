package model

import (
	"errors"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	career "github.com/Erickype/DGC_gestion_indicadores_backend/model/career"
	contractType "github.com/Erickype/DGC_gestion_indicadores_backend/model/contractType"
	dedication "github.com/Erickype/DGC_gestion_indicadores_backend/model/dedication"
	scaledGrade "github.com/Erickype/DGC_gestion_indicadores_backend/model/scaledGrade"
	teacher "github.com/Erickype/DGC_gestion_indicadores_backend/model/teacher"
	"gorm.io/gorm"
	"time"
)

type TeachersList struct {
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	AcademicPeriodID uint      `gorm:"primary_key;" json:"academic_period_id"`
	TeacherID        uint      `gorm:"primary_key;" json:"teacher_id"`
	CareerID         uint      `gorm:"not null;" json:"career_id"`
	DedicationID     uint      `gorm:"not null;" json:"dedication_id"`
	ScaledGradeID    uint      `gorm:"not null;" json:"scaled_grade_id"`
	ContractTypeID   uint      `gorm:"not null;" json:"contract_type_id"`

	AcademicPeriod academicPeriod.AcademicPeriod `json:"-"`
	Teacher        teacher.Teacher               `json:"-"`
	Career         career.Career                 `json:"-"`
	Dedication     dedication.Dedication         `json:"-"`
	ScaledGrade    scaledGrade.ScaledGrade       `json:"-"`
	ContractType   contractType.ContractType     `json:"-"`
}

func (tl *TeachersList) TableName() string {
	return model.IndicatorsInformationSchema + ".teachers_lists"
}

func CreateTeachersList(teachersList *TeachersList) (err error) {
	err = database.DB.Create(teachersList).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("profesor en lista ya existe")
		}
		return err
	}
	return nil
}
