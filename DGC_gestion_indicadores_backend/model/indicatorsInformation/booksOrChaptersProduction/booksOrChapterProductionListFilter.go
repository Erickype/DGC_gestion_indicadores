package model

import (
	"fmt"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	"strings"
)

type FilterBooksOrChaptersProductionListsByAcademicPeriodRequest struct {
	DOI              string `json:"DOI"`
	IsChapter        *bool  `json:"is_chapter"`
	Title            string `json:"title"`
	PublicationDate  string `json:"publication_date"`
	PeerReviewed     *bool  `json:"peer_reviewed"`
	DetailedField    string `json:"detailed_field_id"`
	AcademicPeriodID uint   `json:"academic_period_id"`
	PageSize         int    `json:"page_size"`
	Page             int    `json:"page"`
}

type BooksOrChaptersProductionListsByAcademicPeriodJoined struct {
	ID               uint
	DOI              string `json:"DOI"`
	IsChapter        *bool  `json:"is_chapter"`
	Title            string `json:"title"`
	PublicationDate  string `json:"publication_date"`
	PeerReviewed     *bool  `json:"peer_reviewed"`
	DetailedFieldID  uint   `json:"detailed_field_id"`
	DetailedField    string `json:"detailed_field"`
	AcademicPeriodID uint   `json:"academic_period_id"`
}

type FilterBooksOrChaptersProductionListsByAcademicPeriodResponse struct {
	Count                          int64                                                  `json:"count"`
	PageSize                       int                                                    `json:"page_size"`
	Page                           int                                                    `json:"page"`
	BooksOrChaptersProductionLists []BooksOrChaptersProductionListsByAcademicPeriodJoined `json:"books_or_chapters_production_lists"`
}

func FilterBooksOrChaptersProductionListsByAcademicPeriod(
	booksOrChaptersProductionListsByAcademicPeriodResponse *FilterBooksOrChaptersProductionListsByAcademicPeriodResponse,
	booksOrChaptersProductionListsByAcademicPeriodRequest *FilterBooksOrChaptersProductionListsByAcademicPeriodRequest) (err error) {
	query := database.DB

	var conditions []string
	var values []interface{}

	if booksOrChaptersProductionListsByAcademicPeriodRequest.DOI != "" {
		conditions = append(conditions, "LOWER(bocpl.doi) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(booksOrChaptersProductionListsByAcademicPeriodRequest.DOI)))
	}
	if booksOrChaptersProductionListsByAcademicPeriodRequest.IsChapter != nil {
		conditions = append(conditions, "bocpl.is_chapter = ?")
		values = append(values, fmt.Sprintf("%t", *booksOrChaptersProductionListsByAcademicPeriodRequest.IsChapter))
	}
	if booksOrChaptersProductionListsByAcademicPeriodRequest.Title != "" {
		conditions = append(conditions, "LOWER(bocpl.title) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(booksOrChaptersProductionListsByAcademicPeriodRequest.Title)))
	}
	if booksOrChaptersProductionListsByAcademicPeriodRequest.PublicationDate != "" {
		conditions = append(conditions, "cast(bocpl.publication_date as varchar) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(booksOrChaptersProductionListsByAcademicPeriodRequest.PublicationDate)))
	}
	if booksOrChaptersProductionListsByAcademicPeriodRequest.PeerReviewed != nil {
		conditions = append(conditions, "bocpl.peer_reviewed = ?")
		values = append(values, fmt.Sprintf("%t", *booksOrChaptersProductionListsByAcademicPeriodRequest.PeerReviewed))
	}
	if booksOrChaptersProductionListsByAcademicPeriodRequest.DetailedField != "" {
		conditions = append(conditions, "LOWER(df.name) LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", strings.ToLower(booksOrChaptersProductionListsByAcademicPeriodRequest.DetailedField)))
	}

	if len(conditions) > 0 {
		query = query.Where(strings.Join(conditions, " OR "), values...)
	}

	query = query.Select(
		`bocpl.id,
		bocpl.doi,
		bocpl.is_chapter,
		bocpl.title,
		bocpl.publication_date,
		bocpl.peer_reviewed,
		bocpl.detailed_field_id,
		df.name as detailed_field,
		bocpl.academic_period_id`,
	).
		Joins(`join detailed_fields df on bocpl.detailed_field_id = df.id`).
		Where("bocpl.academic_period_id = ?",
			booksOrChaptersProductionListsByAcademicPeriodRequest.AcademicPeriodID)

	var totalCount int64
	err = query.Table("indicators_information.books_or_chapter_production_lists bocpl").Count(&totalCount).Error
	if err != nil {
		return err
	}

	pageSize := booksOrChaptersProductionListsByAcademicPeriodRequest.PageSize
	page := booksOrChaptersProductionListsByAcademicPeriodRequest.Page
	err = query.
		Order("bocpl.updated_at DESC").
		Scopes(model.Paginate(pageSize, page)).
		Scan(&booksOrChaptersProductionListsByAcademicPeriodResponse.BooksOrChaptersProductionLists).Error
	if err != nil {
		return err
	}

	booksOrChaptersProductionListsByAcademicPeriodResponse.Count = totalCount
	booksOrChaptersProductionListsByAcademicPeriodResponse.PageSize = pageSize
	if page <= 0 {
		booksOrChaptersProductionListsByAcademicPeriodResponse.Page = 1
	} else {
		booksOrChaptersProductionListsByAcademicPeriodResponse.Page = page
	}

	return nil
}
