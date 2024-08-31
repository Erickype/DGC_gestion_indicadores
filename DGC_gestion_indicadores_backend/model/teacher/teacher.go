package model

import (
	"errors"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	person "github.com/Erickype/DGC_gestion_indicadores_backend/model/person"
	"gorm.io/gorm"
	"time"
)

type Teacher struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	PersonID  uint      `gorm:"not null;unique" json:"person_id"`

	Person person.Person `json:"-"`
}

func CreateTeacher(Teacher *Teacher) (err error) {
	err = database.DB.Create(Teacher).Error
	if err != nil {
		if errors.Is(gorm.ErrDuplicatedKey, err) {
			return errors.New("docente ya registrado en el periodo")
		}
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

func GetTeachersByAcademicPeriod(teachers *[]GetByAcademicPeriodResponse, academicPeriodID int) error {
	err := database.DB.Table("teachers").Select(
		`teachers.updated_at, 
		teachers.academic_period_id,
		teachers.person_id, 
		people.name || ' ' || people.lastname as person,
		teachers.career_id, 
		careers.abbreviation as career, 
		teachers.dedication_id, 
		dedications.name as dedication, 
		teachers.scaled_grade_id, 
		scaled_grades.abbreviation as scaled_grade,
		teachers.contract_type_id,
		contract_types.abbreviation as contract_type`,
	).Joins(
		`JOIN people ON people.id = teachers.person_id
		JOIN careers ON careers.id = teachers.career_id
		JOIN dedications ON dedications.id = teachers.dedication_id
		JOIN scaled_grades ON scaled_grades.id = teachers.scaled_grade_id	
		JOIN contract_types ON contract_types.id = teachers.contract_type_id`,
	).Where(
		"teachers.academic_period_id = ?", academicPeriodID,
	).Order(
		"teachers.updated_at DESC",
	).Scan(&teachers).Error

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
