package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	career "github.com/Erickype/DGC_gestion_indicadores_backend/model/career"
	dedication "github.com/Erickype/DGC_gestion_indicadores_backend/model/dedication"
	person "github.com/Erickype/DGC_gestion_indicadores_backend/model/person"
	scaledGrade "github.com/Erickype/DGC_gestion_indicadores_backend/model/scaledGrade"
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	ID               uint `gorm:"primary_key;"`
	AcademicPeriodID uint `gorm:"primary_key;" json:"academic_period_id"`
	PersonID         uint `gorm:"primary_key;" json:"person_id"`
	CareerID         uint `gorm:"not null;" json:"career_id"`
	DedicationID     uint `gorm:"not null;" json:"dedication_id"`
	ScaledGradeID    uint `gorm:"not null;" json:"scaled_grade_id"`

	AcademicPeriod academicPeriod.AcademicPeriod `gorm:"constraint:OnUpdate:CASCADE" json:"-"`
	Person         person.Person                 `gorm:"constraint:OnUpdate:CASCADE" json:"-"`
	Career         career.Career                 `gorm:"constraint:OnUpdate:CASCADE;" json:"-"`
	Dedication     dedication.Dedication         `gorm:"constraint:OnUpdate:CASCADE" json:"-"`
	ScaledGrade    scaledGrade.ScaledGrade       `gorm:"constraint:OnUpdate:CASCADE" json:"-"`
}

func CreateTeacher(Teacher *Teacher) (err error) {
	err = database.DB.Create(Teacher).Error
	if err != nil {
		return err
	}
	return nil
}

func GetTeachers(Teachers *[]Teacher) (err error) {
	err = database.DB.Find(Teachers).Error
	if err != nil {
		return err
	}
	return nil
}

func GetTeacher(Teacher *Teacher, id int) (err error) {
	err = database.DB.Where("id = ?", id).First(Teacher).Error
	if err != nil {
		return err
	}
	return nil
}

func GetTeachersByAcademicPeriod(Teachers *[]GetJoin, academicPeriodID int) (err error) {
	err = database.DB.Table("teachers").Select(
		"teachers.updated_at, teachers.id, teachers.academic_period_id,"+
			"teachers.person_id, people.name || ' ' || people.lastname as person,"+
			"teachers.career_id, careers.abbreviation as career, "+
			"teachers.dedication_id, dedications.name as dedication, "+
			"teachers.scaled_grade_id, scaled_grades.abbreviation as scaled_grade").Joins(
		"join people on people.id = teachers.person_id "+
			"join careers on careers.id = teachers.career_id "+
			"join dedications on dedications.id = teachers.dedication_id "+
			"join scaled_grades on scaled_grades.id = teachers.scaled_grade_id").Where(
		"teachers.academic_period_id = ?", academicPeriodID).Order(
		"teachers.updated_at desc").Scan(&Teachers).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateTeacher(Teacher *Teacher) (err error) {
	err = database.DB.Save(Teacher).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteTeacher(id int) (err error) {
	err = database.DB.Delete(&Teacher{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
