package model

import (
	"fmt"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	"strings"
)

type FilterAuthorsResponse struct {
	Count    int64          `json:"count"`
	PageSize int            `json:"page_size"`
	Page     int            `json:"page"`
	Authors  []AuthorPerson `json:"authors"`
}

type AuthorPerson struct {
	ID       uint
	PersonID uint   `json:"person_id"`
	Identity string `json:"identity,omitempty"`
	Name     string `json:"name,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Email    string `json:"email,omitempty"`
}

type FilterAuthorsRequest struct {
	Identity string `json:"identity,omitempty"`
	Name     string `json:"name,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Email    string `json:"email,omitempty"`
	PageSize int    `json:"page_size"`
	Page     int    `json:"page"`
}

func FilterAuthors(filterTeachersResponse *FilterAuthorsResponse, filterTeachersRequest *FilterAuthorsRequest) (err error) {
	query := database.DB

	var conditions []string
	var values []interface{}

	if filterTeachersRequest.Identity != "" {
		conditions = append(conditions, "LOWER(people.identity) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterTeachersRequest.Identity)))
	}
	if filterTeachersRequest.Name != "" {
		conditions = append(conditions, "LOWER(people.name) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterTeachersRequest.Name)))
	}
	if filterTeachersRequest.Lastname != "" {
		conditions = append(conditions, "LOWER(people.lastname) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterTeachersRequest.Lastname)))
	}
	if filterTeachersRequest.Email != "" {
		conditions = append(conditions, "LOWER(people.email) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterTeachersRequest.Email)))
	}

	if len(conditions) > 0 {
		query = query.Where(strings.Join(conditions, " OR "), values...)
	}

	query = query.Select(
		`authors.id,
		authors.person_id,
		people.identity,
		people.name,
		people.lastname,
		people.email`,
	).Joins(`INNER JOIN people on people.id = authors.person_id`)

	var totalCount int64
	err = query.Model(&Author{}).Count(&totalCount).Error
	if err != nil {
		return err
	}

	pageSize := filterTeachersRequest.PageSize
	page := filterTeachersRequest.Page
	err = query.
		Order("authors.updated_at DESC").
		Scopes(model.Paginate(pageSize, page)).
		Scan(&filterTeachersResponse.Authors).Error
	if err != nil {
		return err
	}

	filterTeachersResponse.Count = totalCount
	filterTeachersResponse.PageSize = pageSize
	if page <= 0 {
		filterTeachersResponse.Page = 1
	} else {
		filterTeachersResponse.Page = page
	}
	return nil
}

func GetAuthorPersonJoinedByAuthorID(authorID int, response *AuthorPerson) (err error) {
	err = database.DB.Model(&Author{}).Select(
		`authors.id,
		authors.person_id,
		people.identity,
		people.name,
		people.lastname,
		people.email`,
	).Where(
		"authors.id = ?", authorID,
	).Joins(
		`INNER JOIN people on people.id = authors.person_id`,
	).Scan(&response).Error

	if err != nil {
		return err
	}
	return nil
}
