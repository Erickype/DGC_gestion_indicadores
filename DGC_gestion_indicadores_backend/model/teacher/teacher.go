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
			return errors.New("docente ya registrado")
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
