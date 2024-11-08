package model

import (
	"fmt"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	"strconv"
	"strings"
)

type FilterPostgraduateProgramsRequest struct {
	Name      string `json:"name,omitempty"`
	StartYear uint   `json:"start_year,omitempty"`
	EndYear   uint   `json:"end_year,omitempty"`
	IsActive  *bool  `json:"is_active,omitempty"`
	PageSize  int    `json:"page_size"`
	Page      int    `json:"page"`
}

type FilterPostGraduateProgramsResponse struct {
	Count                int64                 `json:"count"`
	PageSize             int                   `json:"page_size"`
	Page                 int                   `json:"page"`
	PostgraduatePrograms []PostgraduateProgram `json:"postgraduate_programs"`
}

func FilterPostGraduatePrograms(
	filterPostGraduateProgramsResponse *FilterPostGraduateProgramsResponse,
	filterPostgraduateProgramsRequest *FilterPostgraduateProgramsRequest) (err error) {
	query := database.DB

	var conditions []string
	var values []interface{}

	if filterPostgraduateProgramsRequest.Name != "" {
		conditions = append(conditions, "LOWER(pp.name) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(filterPostgraduateProgramsRequest.Name)))
	}
	if strconv.Itoa(int(filterPostgraduateProgramsRequest.StartYear)) != "" {
		conditions = append(conditions, "cast(pp.start_year as varchar) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strconv.Itoa(int(filterPostgraduateProgramsRequest.StartYear))))
	}
	if strconv.Itoa(int(filterPostgraduateProgramsRequest.EndYear)) != "" {
		conditions = append(conditions, "cast(pp.end_year as varchar) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strconv.Itoa(int(filterPostgraduateProgramsRequest.EndYear))))
	}
	if filterPostgraduateProgramsRequest.IsActive != nil {
		conditions = append(conditions, "pp.is_active = ?")
		values = append(values, fmt.Sprintf("%t", *filterPostgraduateProgramsRequest.IsActive))
	}

	if len(conditions) > 0 {
		query = query.Where(strings.Join(conditions, " OR "), values...)
	}

	query = query.Select(`pp.name,
			pp.start_year,
			pp.end_year,
			pp.is_active`)

	var totalCount int64
	err = query.Table("indicators_information.postgraduate_programs pp").Count(&totalCount).Error
	if err != nil {
		return err
	}

	pageSize := filterPostgraduateProgramsRequest.PageSize
	page := filterPostgraduateProgramsRequest.Page
	err = query.
		Order("pp.start_year desc, pp.is_active desc").
		Scopes(model.Paginate(pageSize, page)).
		Scan(&filterPostGraduateProgramsResponse.PostgraduatePrograms).Error
	if err != nil {
		return err
	}

	filterPostGraduateProgramsResponse.Count = totalCount
	filterPostGraduateProgramsResponse.PageSize = pageSize
	if page <= 0 {
		filterPostGraduateProgramsResponse.Page = 1
	} else {
		filterPostGraduateProgramsResponse.Page = page
	}

	return nil
}
