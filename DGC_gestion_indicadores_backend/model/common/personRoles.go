package model

import (
	"errors"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	author "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicProduction/author"
	person "github.com/Erickype/DGC_gestion_indicadores_backend/model/person"
	teacher "github.com/Erickype/DGC_gestion_indicadores_backend/model/teacher"
	"gorm.io/gorm"
	"slices"
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
		request.Person.Name = strings.ToUpper(request.Person.Name)
		request.Person.Lastname = strings.ToUpper(request.Person.Lastname)
		if err = tx.Create(&request.Person).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return errors.New("persona ya registrada")
			}
			return err
		}
		return InsertPersonRolesFromRolesList(tx, request.Person.ID, request.Roles)
	})
}

func UpdatePersonWithRoles(request *UpdatePersonWithRolesRequest) (err error) {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if len(request.Roles) <= 0 {
			return errors.New("seleccione al menos un rol")
		}

		request.Person.Name = strings.ToUpper(request.Person.Name)
		request.Person.Lastname = strings.ToUpper(request.Person.Lastname)
		err = tx.Save(&request.Person).Error
		if err != nil {
			return err
		}

		var roles []string
		err = GetPersonRoles(int(request.Person.ID), &roles)
		if err != nil {
			return err
		}

		if len(request.Roles) == len(roles) {
			var findIndexes []int
			for _, role := range roles {
				index := slices.IndexFunc(request.Roles, func(s string) bool {
					return s == role
				})
				findIndexes = append(findIndexes, index)
			}
			var notFoundRoles int
			for _, index := range findIndexes {
				if index == -1 {
					notFoundRoles++
				}
			}
			if notFoundRoles <= 0 {
				return nil
			}

			var toDelete = GetRolesToUpdate(request.Roles, roles)
			err = DeletePersonRolesFromRolesLists(tx, request.Person.ID, toDelete)
			if err != nil {
				return err
			}

			var toInsert = GetRolesToUpdate(roles, request.Roles)
			return InsertPersonRolesFromRolesList(tx, request.Person.ID, toInsert)
		}

		//Delete
		if len(request.Roles) < len(roles) {
			var toDelete = GetRolesToUpdate(request.Roles, roles)
			return DeletePersonRolesFromRolesLists(tx, request.Person.ID, toDelete)
		}

		//Insert
		if len(request.Roles) > len(roles) {
			var toInsert = GetRolesToUpdate(roles, request.Roles)
			return InsertPersonRolesFromRolesList(tx, request.Person.ID, toInsert)
		}

		return nil
	})
}

func GetRolesToUpdate(currentRoles, requestRoles []string) (rolesToUpdate []string) {
	var findIndexes []int
	for _, role := range currentRoles {
		index := slices.IndexFunc(requestRoles, func(s string) bool {
			return s == role
		})
		findIndexes = append(findIndexes, index)
	}

	for i, role := range requestRoles {
		if !slices.Contains(findIndexes, i) {
			rolesToUpdate = append(rolesToUpdate, role)
		}
	}
	return rolesToUpdate
}

func DeletePersonRolesFromRolesLists(tx *gorm.DB, personID uint, Roles []string) (err error) {
	for _, role := range Roles {
		switch role {
		case model.PersonRoleTeacher:
			err = tx.Delete(&teacher.Teacher{}, "person_id = ?", personID).Error
			if err != nil {
				if errors.Is(err, gorm.ErrForeignKeyViolated) {
					return errors.New("profesor en uso")
				}
				return err
			}
			break
		case model.PersonRoleAuthor:
			err = tx.Delete(&author.Author{}, "person_id = ?", personID).Error
			if err != nil {
				if errors.Is(err, gorm.ErrForeignKeyViolated) {
					return errors.New("autor en uso")
				}
				return err
			}
			break
		}
	}
	return nil
}

func InsertPersonRolesFromRolesList(tx *gorm.DB, personID uint, Roles []string) (err error) {
	for _, roleID := range Roles {
		switch roleID {
		case model.PersonRoleTeacher:
			newAuthor := teacher.Teacher{
				PersonID: personID,
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
				PersonID: personID,
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
