package model

import (
	"errors"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicProduction "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicProduction/author"
	career "github.com/Erickype/DGC_gestion_indicadores_backend/model/career"
	"gorm.io/gorm"
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

type PostBooksOrChaptersProductionListsAuthorCareersRequest struct {
	BooksOrChaptersProductionListID uint   `json:"books_or_chapters_production_list_id"`
	AuthorID                        uint   `json:"author_id"`
	Careers                         []uint `json:"careers"`
}

func (bc BooksOrChaptersProductionListAuthor) TableName() string {
	return model.IndicatorsInformationSchema + ".books_or_chapter_production_lists_authors"
}

func PostBooksOrChaptersProductionListsAuthorCareers(
	request *PostBooksOrChaptersProductionListsAuthorCareersRequest) (err error) {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		for _, careerID := range request.Careers {
			booksOrChaptersProductionListAuthor := BooksOrChaptersProductionListAuthor{
				BooksOrChaptersProductionListID: request.BooksOrChaptersProductionListID,
				AuthorID:                        request.AuthorID,
				CareerID:                        careerID,
			}
			if err := tx.Create(&booksOrChaptersProductionListAuthor).Error; err != nil {
				if errors.Is(err, gorm.ErrForeignKeyViolated) {
					return errors.New("claves no encontradas")
				}
				if errors.Is(err, gorm.ErrDuplicatedKey) {
					return errors.New("autor ya registrado")
				}
				return err
			}
		}
		return nil
	})
}
