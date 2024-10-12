package model

import (
	"fmt"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	"strings"
)

type FilterSocialProjectListsByAcademicPeriodRequest struct {
	AcademicPeriodID int    `json:"academic_period_id"`
	Career           string `json:"career"`
	Name             string `json:"name"`
	PageSize         int    `json:"page_size"`
	Page             int    `json:"page"`
}

type SocialProjectListByAcademicPeriodJoined struct {
	ID               uint
	AcademicPeriodID uint   `json:"academic_period_id"`
	CareerID         uint   `json:"career_id"`
	Career           string `json:"career"`
	Name             string `json:"name"`
}

type FilterSocialProjectListsByAcademicPeriodResponse struct {
	Count              int64                                     `json:"count"`
	PageSize           int                                       `json:"page_size"`
	Page               int                                       `json:"page"`
	SocialProjectLists []SocialProjectListByAcademicPeriodJoined `json:"social_project_lists"`
}

func FilterSocialProjectListsByAcademicPeriod(
	filterSocialProjectListsByAcademicPeriodResponse *FilterSocialProjectListsByAcademicPeriodResponse,
	filterSocialProjectListsByAcademicPeriodRequest *FilterSocialProjectListsByAcademicPeriodRequest) (err error) {
	query := database.DB

	var conditions []string
	var values []interface{}

	if filterSocialProjectListsByAcademicPeriodRequest.Career != "" {
		conditions = append(conditions, "LOWER(c.name) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%",
			strings.ToLower(filterSocialProjectListsByAcademicPeriodRequest.Career)))
	}
	if filterSocialProjectListsByAcademicPeriodRequest.Name != "" {
		conditions = append(conditions, "LOWER(spl.name) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%",
			strings.ToLower(filterSocialProjectListsByAcademicPeriodRequest.Name)))
	}

	if len(conditions) > 0 {
		query = query.Where(strings.Join(conditions, " OR "), values...)
	}

	query = query.
		Select(`spl.id,
		spl.academic_period_id,
		spl.career_id,
		c.name as career,
		spl.name`,
		).
		Joins(`join careers c on spl.career_id = c.id`).
		Where("spl.academic_period_id = ?",
			filterSocialProjectListsByAcademicPeriodRequest.AcademicPeriodID)

	var totalCount int64
	err = query.Table("indicators_information.social_project_lists spl").Count(&totalCount).Error
	if err != nil {
		return err
	}

	pageSize := filterSocialProjectListsByAcademicPeriodRequest.PageSize
	page := filterSocialProjectListsByAcademicPeriodRequest.Page
	err = query.
		Order("spl.updated_at DESC").
		Scopes(model.Paginate(pageSize, page)).
		Scan(&filterSocialProjectListsByAcademicPeriodResponse.SocialProjectLists).Error
	if err != nil {
		return err
	}

	filterSocialProjectListsByAcademicPeriodResponse.Count = totalCount
	filterSocialProjectListsByAcademicPeriodResponse.PageSize = pageSize
	if page <= 0 {
		filterSocialProjectListsByAcademicPeriodResponse.Page = 1
	} else {
		filterSocialProjectListsByAcademicPeriodResponse.Page = page
	}

	return nil
}
