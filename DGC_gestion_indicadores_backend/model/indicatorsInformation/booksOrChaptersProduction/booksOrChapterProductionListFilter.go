package model

import (
	"gorm.io/datatypes"
)

type FilterBooksOrChaptersProductionListsByAcademicPeriodRequest struct {
	DOI              string         `json:"DOI"`
	IsChapter        *bool          `json:"is_chapter"`
	Title            string         `json:"title"`
	PublicationDate  datatypes.Date `json:"publication_date"`
	PeerReviewed     *bool          `json:"peer_reviewed"`
	AcademicPeriodID uint           `json:"academic_period_id"`
	DetailedFieldID  uint           `json:"detailed_field_id"`
}

type BooksOrChaptersProductionListsByAcademicPeriodJoined struct {
	ID               uint
	DOI              string         `json:"DOI"`
	IsChapter        *bool          `json:"is_chapter"`
	Title            string         `json:"title"`
	PublicationDate  datatypes.Date `json:"publication_date"`
	PeerReviewed     *bool          `json:"peer_reviewed"`
	AcademicPeriodID uint           `json:"academic_period_id"`
	AcademicPeriod   string         `json:"academic_period"`
	DetailedFieldID  uint           `json:"detailed_field_id"`
	DetailedField    string         `json:"detailed_field"`
}

type FilterBooksOrChaptersProductionListsByAcademicPeriodResponse struct {
	Count                          int64                                                  `json:"count"`
	PageSize                       int                                                    `json:"page_size"`
	Page                           int                                                    `json:"page"`
	BooksOrChaptersProductionLists []BooksOrChaptersProductionListsByAcademicPeriodJoined `json:"books_or_chapters_production_lists"`
}
