package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicProduction "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicProduction/author"
	career "github.com/Erickype/DGC_gestion_indicadores_backend/model/career"
	"time"
)

type BooksOrChaptersProductionListAuthor struct {
	CreatedAt                       time.Time `json:"created_at"`
	UpdatedAt                       time.Time `json:"updated_at"`
	BooksOrChaptersProductionListID uint      `gorm:"primary_key" json:"books_or_chapters_production_list_id;"`
	AuthorID                        uint      `gorm:"primary_key" json:"author_id"`
	CareerID                        uint      `gorm:"primary_key" json:"career_id"`

	BooksOrChaptersProductionList BooksOrChaptersProductionList `json:"-"`
	Author                        academicProduction.Author     `json:"-"`
	Career                        career.Career                 `json:"-"`
}

func (bc BooksOrChaptersProductionListAuthor) TableName() string {
	return model.IndicatorsInformationSchema + ".books_or_chapter_production_lists_authors"
}
