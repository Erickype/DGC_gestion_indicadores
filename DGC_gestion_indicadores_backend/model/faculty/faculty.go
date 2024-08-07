package model

import (
	"errors"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"gorm.io/gorm"
)

type Faculty struct {
	gorm.Model
	ID           uint   `gorm:"primary_key"`
	Name         string `gorm:"size:200;not null;unique" json:"name"`
	Description  string `gorm:"size:50;" json:"description"`
	Abbreviation string `gorm:"size:50;not null;unique" json:"abbreviation"`
}

func CreateFaculty(faculty *Faculty) (err error) {
	err = database.DB.Create(faculty).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("facultad ya existe")
		}
		return err
	}
	return nil
}

func GetFaculties(faculties *[]Faculty) (err error) {
	err = database.DB.Order("created_at desc").Find(faculties).Error
	if err != nil {
		return err
	}
	return nil
}

func GetFaculty(faculty *Faculty, id int) (err error) {
	err = database.DB.Where("id = ?", id).First(faculty).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("facultad no encontrada")
		}
		return err
	}
	return nil
}

func UpdateFaculty(faculty *Faculty) (err error) {
	err = database.DB.Save(faculty).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("facultad no existe")
		}
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("facultad ya existe")
		}
		return err
	}
	return nil
}

func DeleteFaculty(id int) (err error) {
	err = database.DB.Unscoped().Delete(&Faculty{}, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("facultad no existe")
		}
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return errors.New("facultad en uso")
		}
	}
	return nil
}
