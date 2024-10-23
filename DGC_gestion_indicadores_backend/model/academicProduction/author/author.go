package model

import (
	"errors"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model/person"
	"gorm.io/gorm"
	"time"
)

type Author struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	PersonID  uint      `gorm:"unique; not null" json:"person_id"`

	Person model.Person `json:"-"`
}

type PostAuthorFromPersonRequest struct {
	Identity string `json:"identity,omitempty"`
	Name     string `json:"name,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Email    string `json:"email,omitempty"`
}

func PostAuthor(request *Author) (err error) {
	err = database.DB.Create(request).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("autor ya registrado")
		}
		return err
	}
	return nil
}
