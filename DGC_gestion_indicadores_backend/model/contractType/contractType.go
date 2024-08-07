package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"time"
)

type ContractType struct {
	ID           uint      `gorm:"primary_key"`
	CreatedAt    time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"not null" json:"updated_at"`
	Name         string    `gorm:"size:200;not null;unique" json:"name"`
	Abbreviation string    `gorm:"size:50;not null;unique" json:"abbreviation"`
	Description  string    `gorm:"size:200;" json:"description,omitempty"`
}

func GetDedications(contractTypes *[]ContractType) (err error) {
	err = database.DB.Find(contractTypes).Error
	if err != nil {
		return err
	}
	return nil
}
