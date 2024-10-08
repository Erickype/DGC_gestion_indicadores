package model

import (
	"fmt"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	"strings"
)

type FilterAcademicProductionListsByAcademicPeriodRequest struct {
	AcademicPeriodID       int    `json:"academic_period_id"`
	DOI                    string `json:"doi,omitempty"`
	PublicationDate        string `json:"publication_date,omitempty"`
	PublicationName        string `json:"publication_name,omitempty"`
	DetailedField          string `json:"detailed_field,omitempty"`
	ScienceMagazine        string `json:"science_magazine,omitempty"`
	ImpactCoefficient      string `json:"impact_coefficient,omitempty"`
	InterculturalComponent *bool  `json:"intercultural_component,omitempty"`
	PageSize               int    `json:"page_size"`
	Page                   int    `json:"page"`
}

type AcademicProductionListByAcademicPeriodJoined struct {
	ID                     uint
	DOI                    string `json:"doi"`
	PublicationDate        string `json:"publication_date"`
	PublicationName        string `json:"publication_name"`
	DetailedFieldID        uint   `json:"detailed_field_id"`
	DetailedField          string `json:"detailed_field"`
	ScienceMagazineID      uint   `json:"science_magazine_id"`
	ScienceMagazine        string `json:"science_magazine"`
	ImpactCoefficientID    uint   `json:"impact_coefficient_id"`
	ImpactCoefficient      string `json:"impact_coefficient"`
	InterculturalComponent *bool  `json:"intercultural_component"`
}

type FilterAcademicProductionListsByAcademicPeriodResponse struct {
	Count                   int64                                          `json:"count"`
	PageSize                int                                            `json:"page_size"`
	Page                    int                                            `json:"page"`
	AcademicProductionLists []AcademicProductionListByAcademicPeriodJoined `json:"academic_production_lists"`
}

func FilterAcademicProductionListsByAcademicPeriod(
	filterAcademicProductionListsResponse *FilterAcademicProductionListsByAcademicPeriodResponse,
	filterAcademicProductionListsRequest *FilterAcademicProductionListsByAcademicPeriodRequest) (err error) {
	query := database.DB

	var conditions []string
	var values []interface{}

	if filterAcademicProductionListsRequest.DOI != "" {
		conditions = append(conditions, "LOWER(apl.doi) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterAcademicProductionListsRequest.DOI)))
	}
	if filterAcademicProductionListsRequest.PublicationDate != "" {
		conditions = append(conditions, "LOWER(apl.publication_date) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterAcademicProductionListsRequest.PublicationDate)))
	}
	if filterAcademicProductionListsRequest.PublicationName != "" {
		conditions = append(conditions, "LOWER(apl.publication_name) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterAcademicProductionListsRequest.PublicationName)))
	}
	if filterAcademicProductionListsRequest.DetailedField != "" {
		conditions = append(conditions, "LOWER(df.name) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterAcademicProductionListsRequest.DetailedField)))
	}
	if filterAcademicProductionListsRequest.ScienceMagazine != "" {
		conditions = append(conditions, "LOWER(sm.name) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterAcademicProductionListsRequest.ScienceMagazine)))
	}
	if filterAcademicProductionListsRequest.ImpactCoefficient != "" {
		conditions = append(conditions, "LOWER(ad.name || ' ' || cf.name) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterAcademicProductionListsRequest.ImpactCoefficient)))
	}
	if filterAcademicProductionListsRequest.InterculturalComponent != nil {
		conditions = append(conditions, "apl.intercultural_component = ?")
		values = append(values, fmt.Sprintf("%%%b%%", filterAcademicProductionListsRequest.InterculturalComponent))
	}

	if len(conditions) > 0 {
		query = query.Where(strings.Join(conditions, " OR "), values...)
	}

	query = query.Select(
		`apl.id,
		apl.doi,
		apl.publication_date,
		apl.publication_name,
		apl.detailed_field_id,
		df.name as detailed_field,
		apl.science_magazine_id,
		sm.name as science_magazine,
		apl.impact_coefficient_id,
		ad.name || ' ' || cf.name as impact_coefficient,
		apl.intercultural_component`,
	).Joins(
		`join science_magazines sm on apl.science_magazine_id = sm.id
		join impact_coefficients ic on apl.impact_coefficient_id = ic.id
		join academic_databases ad on ic.academic_database_id = ad.id
		join compensation_factors cf on ic.compensation_factor_id = cf.id
		join detailed_fields df on apl.detailed_field_id = df.id`,
	).Where(
		"apl.academic_period_id = ?", filterAcademicProductionListsRequest.AcademicPeriodID,
	)

	var totalCount int64
	err = query.Table("indicators_information.academic_production_lists apl").Count(&totalCount).Error
	if err != nil {
		return err
	}

	pageSize := filterAcademicProductionListsRequest.PageSize
	page := filterAcademicProductionListsRequest.Page
	err = query.
		Order("apl.updated_at DESC").
		Scopes(model.Paginate(pageSize, page)).
		Scan(&filterAcademicProductionListsResponse.AcademicProductionLists).Error
	if err != nil {
		return err
	}

	filterAcademicProductionListsResponse.Count = totalCount
	filterAcademicProductionListsResponse.PageSize = pageSize
	if page <= 0 {
		filterAcademicProductionListsResponse.Page = 1
	} else {
		filterAcademicProductionListsResponse.Page = page
	}
	return nil
}
