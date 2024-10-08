package model

import (
	"errors"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicProduction "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicProduction/author"
	career "github.com/Erickype/DGC_gestion_indicadores_backend/model/career"
	indicatorsInformation "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation"
	"gorm.io/gorm"
	"time"
)

type BooksOrChaptersProductionListAuthor struct {
	CreatedAt                       time.Time `json:"created_at"`
	UpdatedAt                       time.Time `json:"updated_at"`
	BooksOrChaptersProductionListID uint      `gorm:"primary_key" json:"books_or_chapters_production_list_id;"`
	AuthorID                        uint      `gorm:"primary_key" json:"author_id"`

	BooksOrChaptersProductionList BooksOrChaptersProductionList `json:"-"`
	Author                        academicProduction.Author     `json:"-"`
}

type PostBooksOrChaptersProductionListsAuthorCareersRequest struct {
	BooksOrChaptersProductionListID uint   `json:"books_or_chapters_production_list_id"`
	AuthorID                        uint   `json:"author_id"`
	Careers                         []uint `json:"careers"`
}

type UpdateBooksOrChaptersProductionListsAuthorCareersRequest struct {
	BooksOrChaptersProductionListID uint   `json:"books_or_chapters_production_list_id"`
	AuthorID                        uint   `json:"author_id"`
	Careers                         []uint `json:"careers"`
}

type BooksOrChaptersProductionListsAuthorsCareersJoined struct {
	AuthorID uint            `json:"author_id"`
	Author   string          `json:"author"`
	Careers  []career.Career `json:"careers,omitempty"`
}

func (bc BooksOrChaptersProductionListAuthor) TableName() string {
	return model.IndicatorsInformationSchema + ".books_or_chapter_production_lists_authors"
}

func PostBooksOrChaptersProductionListsAuthorCareers(
	request *PostBooksOrChaptersProductionListsAuthorCareersRequest) (err error) {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		booksOrChaptersProductionListAuthor := BooksOrChaptersProductionListAuthor{
			BooksOrChaptersProductionListID: request.BooksOrChaptersProductionListID,
			AuthorID:                        request.AuthorID,
		}
		if err = tx.Create(&booksOrChaptersProductionListAuthor).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return errors.New("autor ya registrado en el libro o capítulo")
			}
			return err
		}

		var academicPeriodID uint
		err = database.DB.Table("indicators_information.books_or_chapter_production_lists bocpl").
			Select("bocpl.academic_period_id").
			Where("bocpl.id = ?", request.BooksOrChaptersProductionListID).
			Scan(&academicPeriodID).Error
		if err != nil {
			return err
		}
		if academicPeriodID == 0 {
			return errors.New("no se encontró periodo académico")
		}

		err = database.DB.Delete(
			&indicatorsInformation.AcademicPeriodAuthorCareer{},
			"academic_period_id = ? and author_id = ?",
			academicPeriodID, request.AuthorID).
			Error
		if err != nil {
			return err
		}

		for _, careerID := range request.Careers {
			academicPeriodAuthorCareer := indicatorsInformation.AcademicPeriodAuthorCareer{
				AcademicPeriodID: academicPeriodID,
				AuthorID:         request.AuthorID,
				CareerID:         careerID,
			}
			if err = tx.Create(&academicPeriodAuthorCareer).Error; err != nil {
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

func UpdateBooksOrChaptersProductionListsAuthorCareers(
	request *UpdateBooksOrChaptersProductionListsAuthorCareersRequest) (err error) {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		var academicPeriodID uint
		err = database.DB.Table("indicators_information.books_or_chapter_production_lists bocpl").
			Select("bocpl.academic_period_id").
			Where("bocpl.id = ?", request.BooksOrChaptersProductionListID).
			Scan(&academicPeriodID).Error
		if err != nil {
			return err
		}
		if academicPeriodID == 0 {
			return errors.New("no se encontró periodo académico")
		}

		err = database.DB.Delete(
			&indicatorsInformation.AcademicPeriodAuthorCareer{},
			"academic_period_id = ? and author_id = ?",
			academicPeriodID, request.AuthorID).
			Error
		if err != nil {
			return err
		}

		for _, careerID := range request.Careers {
			academicPeriodAuthorCareer := indicatorsInformation.AcademicPeriodAuthorCareer{
				AcademicPeriodID: academicPeriodID,
				AuthorID:         request.AuthorID,
				CareerID:         careerID,
			}
			if err = tx.Create(&academicPeriodAuthorCareer).Error; err != nil {
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

func GetBooksOrChaptersProductionListsAuthorsJoinedByBooksOrChaptersProductionListID(
	academicProductionListID int, response *[]BooksOrChaptersProductionListsAuthorsCareersJoined) (err error) {
	var academicProductionAuthors []struct {
		Author   string `json:"author"`
		AuthorId uint   `json:"author_id"`
	}
	err = database.DB.Table("indicators_information.books_or_chapter_production_lists_authors bocpla").
		Select(`p.name || ' ' || p.lastname as author,
				bocpla.author_id`).
		Joins("join authors a on bocpla.author_id = a.id").
		Joins("join people p on a.person_id = p.id").
		Group("bocpla.author_id, p.name, p.lastname").
		Where("bocpla.books_or_chapters_production_list_id = ?", academicProductionListID).
		Scan(&academicProductionAuthors).Error
	if err != nil {
		return err
	}

	for _, authorCareer := range academicProductionAuthors {
		var authorCareers []career.Career
		err = database.DB.Table("indicators_information.books_or_chapter_production_lists_authors bocpla").
			Select("c.*").
			Joins(`join indicators_information.books_or_chapter_production_lists bocpl 
    					on bocpla.books_or_chapters_production_list_id = bocpl.id`).
			Joins(`join indicators_information.academic_period_author_careers apac
    					on bocpla.author_id = apac.author_id and bocpl.academic_period_id = apac.academic_period_id`).
			Joins("join careers c on apac.career_id = c.id").
			Where("bocpla.author_id = ?", authorCareer.AuthorId).
			Where("bocpla.books_or_chapters_production_list_id = ?", academicProductionListID).
			Scan(&authorCareers).Error
		if err != nil {
			return err
		}

		authorCareersJoined := BooksOrChaptersProductionListsAuthorsCareersJoined{
			AuthorID: authorCareer.AuthorId,
			Author:   authorCareer.Author,
			Careers:  authorCareers,
		}
		*response = append(*response, authorCareersJoined)
	}
	return nil
}
