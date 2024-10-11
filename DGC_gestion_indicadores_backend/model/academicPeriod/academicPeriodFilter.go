package model

import (
	"fmt"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	"strings"
)

type FilterAcademicPeriodsRequest struct {
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	Abbreviation string `json:"abbreviation,omitempty"`
	StartDate    string `json:"start_date,omitempty"`
	EndDate      string `json:"end_date,omitempty"`
	PageSize     int    `json:"page_size"`
	Page         int    `json:"page"`
}

type FilterAcademicPeriodsResponse struct {
	Count           int64            `json:"count"`
	PageSize        int              `json:"page_size"`
	Page            int              `json:"page"`
	AcademicPeriods []AcademicPeriod `json:"academic_periods"`
}

func FilterAcademicPeriods(
	filterAcademicPeriodsResponse *FilterAcademicPeriodsResponse,
	filterAcademicPeriodsRequest *FilterAcademicPeriodsRequest) (err error) {
	query := database.DB

	var conditions []string
	var values []interface{}

	if filterAcademicPeriodsRequest.Name != "" {
		conditions = append(conditions, "LOWER(ap.name) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterAcademicPeriodsRequest.Name)))
	}
	if filterAcademicPeriodsRequest.Description != "" {
		conditions = append(conditions, "LOWER(ap.description) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterAcademicPeriodsRequest.Description)))
	}
	if filterAcademicPeriodsRequest.Abbreviation != "" {
		conditions = append(conditions, "LOWER(ap.abbreviation) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterAcademicPeriodsRequest.Abbreviation)))
	}
	if filterAcademicPeriodsRequest.StartDate != "" {
		conditions = append(conditions, "cast(ap.start_date as varchar) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterAcademicPeriodsRequest.StartDate)))
	}
	if filterAcademicPeriodsRequest.EndDate != "" {
		conditions = append(conditions, "cast(ap.end_date as varchar) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterAcademicPeriodsRequest.EndDate)))
	}

	if len(conditions) > 0 {
		query = query.Where(strings.Join(conditions, " OR "), values...)
	}

	query = query.Select(`ap.id,
			ap.name,
			ap.description,
			ap.abbreviation,
			ap.start_date,
			ap.end_date`)

	var totalCount int64
	err = query.Table("academic_periods ap").Count(&totalCount).Error
	if err != nil {
		return err
	}

	pageSize := filterAcademicPeriodsRequest.PageSize
	page := filterAcademicPeriodsRequest.Page
	err = query.
		Order("ap.start_date DESC").
		Scopes(model.Paginate(pageSize, page)).
		Scan(&filterAcademicPeriodsResponse.AcademicPeriods).Error
	if err != nil {
		return err
	}

	filterAcademicPeriodsResponse.Count = totalCount
	filterAcademicPeriodsResponse.PageSize = pageSize
	if page <= 0 {
		filterAcademicPeriodsResponse.Page = 1
	} else {
		filterAcademicPeriodsResponse.Page = page
	}
	return nil
}
