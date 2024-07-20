package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	Identity string `gorm:"size:13;not null;unique" json:"identity"`
	Name     string `gorm:"size:50;not null;" json:"name"`
	Lastname string `gorm:"size:50;not null;" json:"lastname"`
	Email    string `gorm:"size:50;not null;" json:"email"`
}

func CreatePerson(person *Person) (err error) {
	err = database.DB.Create(person).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPersons(persons *[]Person) (err error) {
	err = database.DB.Find(persons).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPerson(person *Person, id int) (err error) {
	err = database.DB.Where("id = ?", id).First(person).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdatePerson(person *Person) (err error) {
	err = database.DB.Save(person).Error
	if err != nil {
		return err
	}
	return nil
}

func DeletePerson(id int) (err error) {
	err = database.DB.Delete(&Person{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
