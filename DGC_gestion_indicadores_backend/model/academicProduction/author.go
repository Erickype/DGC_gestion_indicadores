package model

import (
	person "github.com/Erickype/DGC_gestion_indicadores_backend/model/person"
	"time"
)

type Author struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	PersonID  uint      `gorm:"not null" json:"person_id"`

	Person person.Person `json:"-"`
}
