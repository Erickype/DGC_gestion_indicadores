package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	"time"
)

type IndicatorType struct {
	ID           uint      `gorm:"primary_key"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `gorm:"size:200;not null;unique" json:"name"`
	Abbreviation string    `gorm:"size:100;not null;unique" json:"abbreviation"`
	Description  string    `gorm:"size:200;" json:"description,omitempty"`
}

func (it IndicatorType) TableName() string {
	return model.IndicatorsSchema + ".indicator_types"
}
