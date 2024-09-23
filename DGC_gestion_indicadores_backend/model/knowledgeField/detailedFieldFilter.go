package model

import (
	"fmt"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	"strings"
)

type FilterDetailedFieldRequest struct {
	WideField     string `json:"wide_field,omitempty"`
	SpecificField string `json:"specific_field,omitempty"`
	DetailedField string `json:"detailed_field,omitempty"`
	PageSize      int    `json:"page_size"`
	Page          int    `json:"page"`
}

type DetailedFieldJoined struct {
	WideFieldID     uint   `json:"wide_field_id"`
	WideField       string `json:"wide_field"`
	SpecificFieldID uint   `json:"specific_field_id"`
	SpecificField   string `json:"specific_field"`
	DetailedFieldID uint   `json:"detailed_field_id"`
	DetailedField   string `json:"detailed_field"`
}

type FilterDetailedFieldResponse struct {
	Count          int64                 `json:"count"`
	PageSize       int                   `json:"page_size"`
	Page           int                   `json:"page"`
	DetailedFields []DetailedFieldJoined `json:"detailed_fields"`
}

func FilterDetailedFields(
	filterDetailedFieldResponse *FilterDetailedFieldResponse,
	filterDetailedFieldRequest *FilterDetailedFieldRequest) (err error) {
	query := database.DB

	var conditions []string
	var values []interface{}

	if filterDetailedFieldRequest.WideField != "" {
		conditions = append(conditions, "LOWER(wf.name) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterDetailedFieldRequest.WideField)))
	}
	if filterDetailedFieldRequest.SpecificField != "" {
		conditions = append(conditions, "LOWER(sf.name) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterDetailedFieldRequest.SpecificField)))
	}
	if filterDetailedFieldRequest.DetailedField != "" {
		conditions = append(conditions, "LOWER(df.name) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterDetailedFieldRequest.DetailedField)))
	}

	if len(conditions) > 0 {
		query = query.Where(strings.Join(conditions, " OR "), values...)
	}

	query = query.Select(
		`wf.id as wide_field_id,
		wf.name as wide_field,
		sf.id as specific_field_id,
		sf.name as specific_field,
		df.id as detailed_field_id,
		df.name as detailed_field`,
	).
		Joins("join specific_fields sf on sf.id = df.specific_field_id").
		Joins("join wide_fields wf on wf.id = sf.wide_field_id")

	var totalCount int64
	err = query.Table("detailed_fields df").Count(&totalCount).Error
	if err != nil {
		return err
	}

	pageSize := filterDetailedFieldRequest.PageSize
	page := filterDetailedFieldRequest.Page
	err = query.
		Order("df.updated_at DESC").
		Scopes(model.Paginate(pageSize, page)).
		Scan(&filterDetailedFieldResponse.DetailedFields).Error
	if err != nil {
		return err
	}

	filterDetailedFieldResponse.Count = totalCount
	filterDetailedFieldResponse.PageSize = pageSize
	if page <= 0 {
		filterDetailedFieldResponse.Page = 1
	} else {
		filterDetailedFieldResponse.Page = page
	}
	return nil
}

func GetDetailedFilterJoinedByDetailedFieldID(detailedFieldID int, response *DetailedFieldJoined) (err error) {
	err = database.DB.Table("detailed_fields df").
		Select(
			`wf.id as wide_field_id,
			wf.name as wide_field,
			sf.id as specific_field_id,
			sf.name as specific_field,
			df.id as detailed_field_id,
			df.name as detailed_field`,
		).
		Joins(`join specific_fields sf on sf.id = df.specific_field_id`).
		Joins("join wide_fields wf on wf.id = sf.wide_field_id").
		Where("df.id = ?", detailedFieldID).
		Scan(&response).Error

	if err != nil {
		return err
	}
	return nil
}
