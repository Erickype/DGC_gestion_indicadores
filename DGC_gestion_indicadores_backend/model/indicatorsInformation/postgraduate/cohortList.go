package model

import (
	"errors"
	"fmt"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	"gorm.io/gorm"
	"strings"
	"time"
)

type CohortList struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Year      uint      `gorm:"primary_key" json:"year"`
}

type FilterCohortListsRequest struct {
	StartYear int `json:"start_year"`
	EndYear   int `json:"end_year"`
	PageSize  int `json:"page_size"`
	Page      int `json:"page"`
}

type FilterCohortListsResponse struct {
	Count       int64        `json:"count"`
	PageSize    int          `json:"page_size"`
	Page        int          `json:"page"`
	CohortLists []CohortList `json:"cohort_lists"`
}

func (c CohortList) TableName() string {
	return model.IndicatorsInformationSchema + ".cohort_lists"
}

func PostCohortList(cohortList *CohortList) (err error) {
	err = database.DB.Create(cohortList).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("año de cohorte ya existe")
		}
		return err
	}
	return nil
}

func PostFilterCohortLists(
	filterCohortListsResponse *FilterCohortListsResponse,
	filterCohortListsRequest *FilterCohortListsRequest) (err error) {
	query := database.DB

	var conditions []string
	var values []interface{}

	if filterCohortListsRequest.StartYear > 0 {
		conditions = append(conditions, "? <= cl.year")
		values = append(values, fmt.Sprintf("%%%d%%", filterCohortListsRequest.StartYear))
	}
	if filterCohortListsRequest.EndYear > 0 {
		conditions = append(conditions, "cl.year <= ?")
		values = append(values, fmt.Sprintf("%%%d%%", filterCohortListsRequest.EndYear))
	}

	if len(conditions) > 0 {
		query = query.Where(strings.Join(conditions, " AND "), values...)
	}

	query = query.Select(`cl.year`)

	var totalCount int64
	err = query.Table("indicators_information.cohort_lists cl").Count(&totalCount).Error
	if err != nil {
		return err
	}

	pageSize := filterCohortListsRequest.PageSize
	page := filterCohortListsRequest.Page
	err = query.
		Order("cl.year DESC").
		Scopes(model.Paginate(pageSize, page)).
		Scan(&filterCohortListsResponse.CohortLists).Error
	if err != nil {
		return err
	}

	filterCohortListsResponse.Count = totalCount
	filterCohortListsResponse.PageSize = pageSize
	if page <= 0 {
		filterCohortListsResponse.Page = 1
	} else {
		filterCohortListsResponse.Page = page
	}
	return nil
}
