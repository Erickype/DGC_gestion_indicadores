package model

import (
	"fmt"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	"strings"
)

type FilterTeachersListsByAcademicPeriodRequest struct {
	AcademicPeriodID int    `json:"academic_period_id,omitempty"`
	TeacherIdentity  string `json:"teacher_identity,omitempty"`
	TeacherName      string `json:"teacher_name,omitempty"`
	TeacherLastname  string `json:"teacher_lastname,omitempty"`
	Career           string `json:"career,omitempty"`
	Dedication       string `json:"dedication,omitempty"`
	ScaledGrade      string `json:"scaled_grade,omitempty"`
	ContractType     string `json:"contract_type,omitempty"`
	PageSize         int    `json:"page_size"`
	Page             int    `json:"page"`
}

type TeachersListByAcademicPeriodJoined struct {
	TeacherID       uint   `json:"teacher_id"`
	TeacherIdentity string `json:"teacher_identity"`
	Teacher         string `json:"teacher"`
	CareerID        uint   `json:"career_id"`
	Career          string `json:"career"`
	DedicationID    uint   `json:"dedication_id"`
	Dedication      string `json:"dedication"`
	ScaledGradeID   uint   `json:"scaled_grade_id"`
	ScaledGrade     string `json:"scaled_grade"`
	ContractTypeID  uint   `json:"contract_type_id"`
	ContractType    string `json:"contract_type"`
}

type FilterTeachersListsByAcademicPeriodResponse struct {
	Count         int64                                `json:"count"`
	PageSize      int                                  `json:"page_size"`
	Page          int                                  `json:"page"`
	TeachersLists []TeachersListByAcademicPeriodJoined `json:"teachers_lists"`
}

func FilterTeachersListsByAcademicPeriod(
	filterTeachersListsResponse *FilterTeachersListsByAcademicPeriodResponse,
	filterTeachersListsRequest *FilterTeachersListsByAcademicPeriodRequest) (err error) {
	query := database.DB

	var conditions []string
	var values []interface{}

	if filterTeachersListsRequest.TeacherIdentity != "" {
		conditions = append(conditions, "LOWER(people.identity) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterTeachersListsRequest.TeacherIdentity)))
	}
	if filterTeachersListsRequest.TeacherName != "" {
		conditions = append(conditions, "LOWER(people.name) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterTeachersListsRequest.TeacherName)))
	}
	if filterTeachersListsRequest.TeacherLastname != "" {
		conditions = append(conditions, "LOWER(people.lastname) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterTeachersListsRequest.TeacherLastname)))
	}
	if filterTeachersListsRequest.Career != "" {
		conditions = append(conditions, "LOWER(careers.abbreviation) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterTeachersListsRequest.Career)))
	}
	if filterTeachersListsRequest.Dedication != "" {
		conditions = append(conditions, "LOWER(dedications.abbreviation) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterTeachersListsRequest.Career)))
	}
	if filterTeachersListsRequest.ScaledGrade != "" {
		conditions = append(conditions, "LOWER(scaled_grades.abbreviation) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterTeachersListsRequest.Career)))
	}
	if filterTeachersListsRequest.ContractType != "" {
		conditions = append(conditions, "LOWER(contract_types.abbreviation) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterTeachersListsRequest.Career)))
	}

	if len(conditions) > 0 {
		query = query.Where(strings.Join(conditions, " OR "), values...)
	}

	query = query.Select(
		`teachers_lists.teacher_id, 
		people.identity as teacher_identity, 
		people.name || ' ' || people.lastname as teacher,
		teachers_lists.career_id, 
		careers.abbreviation as career, 
		teachers_lists.dedication_id, 
		dedications.name as dedication, 
		teachers_lists.scaled_grade_id, 
		scaled_grades.abbreviation as scaled_grade,
		teachers_lists.contract_type_id,
		contract_types.abbreviation as contract_type`,
	).Joins(
		`JOIN teachers on teachers.id = teachers_lists.teacher_id
		JOIN people ON people.id = teachers.person_id
		JOIN careers ON careers.id = teachers_lists.career_id
		JOIN dedications ON dedications.id = teachers_lists.dedication_id
		JOIN scaled_grades ON scaled_grades.id = teachers_lists.scaled_grade_id	
		JOIN contract_types ON contract_types.id = teachers_lists.contract_type_id`,
	).Where(
		"teachers_lists.academic_period_id = ?", filterTeachersListsRequest.AcademicPeriodID,
	)

	var totalCount int64
	err = query.Model(&TeachersList{}).Count(&totalCount).Error
	if err != nil {
		return err
	}

	pageSize := filterTeachersListsRequest.PageSize
	page := filterTeachersListsRequest.Page
	err = query.
		Order("teachers_lists.updated_at DESC").
		Scopes(model.Paginate(pageSize, page)).
		Scan(&filterTeachersListsResponse.TeachersLists).Error
	if err != nil {
		return err
	}

	filterTeachersListsResponse.Count = totalCount
	filterTeachersListsResponse.PageSize = pageSize
	if page <= 0 {
		filterTeachersListsResponse.Page = 1
	} else {
		filterTeachersListsResponse.Page = page
	}
	return nil
}
