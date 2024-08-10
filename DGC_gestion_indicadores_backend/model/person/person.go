package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"time"
)

type Person struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Identity  string    `gorm:"size:13;not null;unique" json:"identity" json:"identity,omitempty"`
	Name      string    `gorm:"size:50;not null;" json:"name" json:"name,omitempty"`
	Lastname  string    `gorm:"size:50;not null;" json:"lastname" json:"lastname,omitempty"`
	Email     string    `gorm:"size:50;not null;" json:"email" json:"email,omitempty"`
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
