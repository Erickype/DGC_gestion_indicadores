package model

import (
	"fmt"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	"strings"
)

type FilterPeopleResponse struct {
	Count    int64    `json:"count"`
	PageSize int      `json:"page_size"`
	Page     int      `json:"page"`
	People   []Person `json:"people"`
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

	var totalCount int64
	err = query.Model(&Person{}).Count(&totalCount).Error
	if err != nil {
		return err
	}

	pageSize := filterPeopleRequest.PageSize
	page := filterPeopleRequest.Page
	err = query.
		Scopes(model.Paginate(pageSize, page)).
		Find(&filterPeopleResponse.People).Error
	if err != nil {
		return err
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
