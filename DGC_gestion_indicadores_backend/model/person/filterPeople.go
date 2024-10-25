package model

import (
	"fmt"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	"strings"
)

type FilterPeopleResponse struct {
	Count    int64             `json:"count"`
	PageSize int               `json:"page_size"`
	Page     int               `json:"page"`
	People   []PersonWithRoles `json:"people"`
}

type PersonWithRoles struct {
	ID       int
	Identity string   `json:"identity"`
	Name     string   `json:"name"`
	Lastname string   `json:"lastname"`
	Email    string   `json:"email"`
	Roles    []string `json:"roles"`
}

type FilterPeopleRequest struct {
	Identity string `json:"identity,omitempty"`
	Name     string `json:"name,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Email    string `json:"email,omitempty"`
	PageSize int    `json:"page_size"`
	Page     int    `json:"page"`
}

func FilterPeople(filterPeopleResponse *FilterPeopleResponse, filterPeopleRequest *FilterPeopleRequest) (err error) {
	query := database.DB

	var conditions []string
	var values []interface{}

	if filterPeopleRequest.Identity != "" {
		conditions = append(conditions, "LOWER(identity) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterPeopleRequest.Identity)))
	}
	if filterPeopleRequest.Name != "" {
		conditions = append(conditions, "LOWER(name) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterPeopleRequest.Name)))
	}
	if filterPeopleRequest.Lastname != "" {
		conditions = append(conditions, "LOWER(lastname) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterPeopleRequest.Lastname)))
	}
	if filterPeopleRequest.Email != "" {
		conditions = append(conditions, "LOWER(email) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterPeopleRequest.Email)))
	}

	if len(conditions) > 0 {
		query = query.Where(strings.Join(conditions, " OR "), values...)
	}

	query = query.Select(`p.id,
		p.identity,
		p.name,
		p.lastname,
		p.email`)

	var totalCount int64
	err = query.Table("people as p").Count(&totalCount).Error
	if err != nil {
		return err
	}

	var peopleAux []struct {
		ID       int
		Identity string `json:"identity"`
		Name     string `json:"name"`
		Lastname string `json:"lastname"`
		Email    string `json:"email"`
	}
	pageSize := filterPeopleRequest.PageSize
	page := filterPeopleRequest.Page
	err = query.
		Order("updated_at DESC").
		Scopes(model.Paginate(pageSize, page)).
		Find(&peopleAux).Error
	if err != nil {
		return err
	}
	if len(peopleAux) <= 0 {
		filterPeopleResponse.People = make([]PersonWithRoles, 0)
		return nil
	}

	for _, personAux := range peopleAux {
		personWithRoles := PersonWithRoles{
			ID:       personAux.ID,
			Identity: personAux.Identity,
			Name:     personAux.Name,
			Lastname: personAux.Lastname,
			Email:    personAux.Email,
			Roles:    nil,
		}
		filterPeopleResponse.People = append(filterPeopleResponse.People, personWithRoles)
	}

	for i := range filterPeopleResponse.People {
		var rawRoles string

		// Use raw SQL for the role part since GORM does not support array subqueries directly.
		query := `
        SELECT array_agg(role)
        FROM (
            SELECT 'teacher' AS role
            WHERE EXISTS (SELECT 1 FROM teachers t WHERE t.person_id = ?)
            UNION ALL
            SELECT 'author' AS role
            WHERE EXISTS (SELECT 1 FROM authors a WHERE a.person_id = ?)
        ) roles;
    `

		// Use db.Raw to execute the custom SQL query for each person
		err = database.DB.Raw(query, filterPeopleResponse.People[i].ID, filterPeopleResponse.People[i].ID).
			Scan(&rawRoles).Error
		if err != nil {
			return err
		}

		if len(rawRoles) <= 0 {
			filterPeopleResponse.People[i].Roles = []string{}
			return
		}

		roles := strings.Trim(rawRoles, "{}")
		if roles != "" {
			filterPeopleResponse.People[i].Roles = strings.Split(roles, ",")
		} else {
			filterPeopleResponse.People[i].Roles = []string{}
		}
	}

	filterPeopleResponse.Count = totalCount
	filterPeopleResponse.PageSize = pageSize
	if page <= 0 {
		filterPeopleResponse.Page = 1
	} else {
		filterPeopleResponse.Page = page
	}
	return nil
}
