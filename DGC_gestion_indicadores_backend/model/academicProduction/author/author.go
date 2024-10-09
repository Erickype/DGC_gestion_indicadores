package model

import (
	"errors"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	person "github.com/Erickype/DGC_gestion_indicadores_backend/model/person"
	"gorm.io/gorm"
	"time"
)

type Author struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	PersonID  uint      `gorm:"unique; not null" json:"person_id"`

	Person person.Person `json:"-"`
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

func PostAuthorFromPerson(request *PostAuthorFromPersonRequest, response *Author) (err error) {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		newPerson := person.Person{
			Identity: request.Identity,
			Name:     request.Name,
			Lastname: request.Lastname,
			Email:    request.Email,
		}
		if err := tx.Create(&newPerson).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return errors.New("persona ya registrada")
			}
			return err
		}
		author := Author{
			PersonID: newPerson.ID,
		}
		if err := tx.Create(&author).Error; err != nil {
			if errors.Is(err, gorm.ErrForeignKeyViolated) {
				return errors.New("persona no encontrada")
			}
			return err
		}
		return nil
	})
}
