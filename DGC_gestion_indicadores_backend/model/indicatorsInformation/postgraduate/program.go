package model

import (
	"errors"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	"gorm.io/gorm"
	"time"
)

type PostgraduateProgram struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ID        uint      `gorm:"primary_key"`
	Name      string    `gorm:"size:1000;not null;unique" json:"name"`
	StartYear uint      `gorm:"not null" json:"start_year"`
	EndYear   uint      `gorm:"not null" json:"end_year"`
	IsActive  *bool     `json:"is_active"`
}

func (grl PostgraduateProgram) TableName() string {
	return model.IndicatorsInformationSchema + ".postgraduate_programs"
}

func PostPostgraduateProgram(program *PostgraduateProgram) (err error) {
	err = database.DB.Create(program).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("posgrado ya existe")
		}
		return err
	}
	return nil
}
