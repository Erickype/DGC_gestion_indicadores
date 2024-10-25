package model

import (
	"errors"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	author "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicProduction/author"
	person "github.com/Erickype/DGC_gestion_indicadores_backend/model/person"
	teacher "github.com/Erickype/DGC_gestion_indicadores_backend/model/teacher"
	"gorm.io/gorm"
	"strings"
)

type PostPersonWithRolesRequest struct {
	Roles  []string      `json:"roles"`
	Person person.Person `json:"person"`
}

type UpdatePersonWithRolesRequest struct {
	Roles  []string      `json:"roles"`
	Person person.Person `json:"person"`
}

func PostPersonWithRoles(request *PostPersonWithRolesRequest) (err error) {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&request.Person).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return errors.New("persona ya registrada")
			}
			return err
		}
		for _, roleID := range request.Roles {
			switch roleID {
			case model.PersonRoleTeacher:
				newAuthor := teacher.Teacher{
					PersonID: request.Person.ID,
				}
				if err = tx.Create(&newAuthor).Error; err != nil {
					if errors.Is(err, gorm.ErrForeignKeyViolated) {
						return errors.New("persona no encontrada")
					}
					if errors.Is(err, gorm.ErrDuplicatedKey) {
						return errors.New("persona ya registrada como docente")
					}
					return err
				}
				break

			case model.PersonRoleAuthor:
				newTeacher := author.Author{
					PersonID: request.Person.ID,
				}
				if err = tx.Create(&newTeacher).Error; err != nil {
					if errors.Is(err, gorm.ErrForeignKeyViolated) {
						return errors.New("persona no encontrada")
					}
					if errors.Is(err, gorm.ErrDuplicatedKey) {
						return errors.New("persona ya registrada como autor")
					}
					return err
				}
				break
			}
		}
		return nil
	})
}

func UpdatePersonWithRoles(request *UpdatePersonWithRolesRequest) (err error) {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		var roles []string
		err = GetPersonRoles(int(request.Person.ID), &roles)
		if err != nil {
			return err
		}

		return nil
	})
}

func GetPersonRoles(personID int, roles *[]string) (err error) {
	var rawRoles string

	rolesQuery := `
        SELECT array_agg(role)
        FROM (
            SELECT 'teacher' AS role
            WHERE EXISTS (SELECT 1 FROM teachers t WHERE t.person_id = ?)
            UNION ALL
            SELECT 'author' AS role
            WHERE EXISTS (SELECT 1 FROM authors a WHERE a.person_id = ?)
        ) roles;
    `

	err = database.DB.Raw(rolesQuery, personID, personID).
		Scan(&rawRoles).Error
	if err != nil {
		return err
	}

	if len(rawRoles) <= 0 {
		*roles = []string{}
		return
	}

	trimmedRoles := strings.Trim(rawRoles, "{}")
	if trimmedRoles != "" {
		*roles = strings.Split(trimmedRoles, ",")
	} else {
		*roles = []string{}
	}

	return nil
}
