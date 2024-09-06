package model

import (
	"fmt"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	"strings"
)

type FilterTeachersResponse struct {
	Count    int64           `json:"count"`
	PageSize int             `json:"page_size"`
	Page     int             `json:"page"`
	Teachers []TeacherPerson `json:"teachers"`
}

type TeacherPerson struct {
	ID       uint
	PersonID uint   `json:"person_id"`
	Identity string `json:"identity,omitempty"`
	Name     string `json:"name,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Email    string `json:"email,omitempty"`
}

type FilterTeachersRequest struct {
	Identity string `json:"identity,omitempty"`
	Name     string `json:"name,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Email    string `json:"email,omitempty"`
	PageSize int    `json:"page_size"`
	Page     int    `json:"page"`
}

func FilterTeachers(filterTeachersResponse *FilterTeachersResponse, filterTeachersRequest *FilterTeachersRequest) (err error) {
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
		`teachers.id,
		teachers.person_id,
		people.identity,
		people.name,
		people.lastname,
		people.email`,
	).Joins(`INNER JOIN people on people.id = teachers.person_id`)

	var totalCount int64
	err = query.Model(&Teacher{}).Count(&totalCount).Error
	if err != nil {
		return err
	}

	pageSize := filterTeachersRequest.PageSize
	page := filterTeachersRequest.Page
	err = query.
		Order("teachers.updated_at DESC").
		Scopes(model.Paginate(pageSize, page)).
		Scan(&filterTeachersResponse.Teachers).Error
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

func GetTeacherPersonJoinedByTeacherID(teacherID int, response *TeacherPerson) (err error) {
	err = database.DB.Model(&Teacher{}).Select(
		`teachers.id,
		teachers.person_id,
		people.identity,
		people.name,
		people.lastname,
		people.email`,
	).Where(
		"teachers.ID = ?", teacherID,
	).Joins(
		`INNER JOIN people on people.id = teachers.person_id`,
	).Scan(&response).Error

	if err != nil {
		return err
	}
	return nil
}
