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

type AcademicProductionListsAuthor struct {
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
	AcademicProductionListID uint      `gorm:"primary_key" json:"academic_production_list_id"`
	AuthorID                 uint      `gorm:"primary_key" json:"author_id"`

	AcademicProductionList AcademicProductionList    `json:"-"`
	Author                 academicProduction.Author `json:"-"`
}

type AcademicProductionListsAuthorsCareersJoined struct {
	AuthorID uint            `json:"author_id"`
	Author   string          `json:"author"`
	Careers  []career.Career `json:"careers,omitempty"`
}

type PostAcademicProductionListsAuthorCareersRequest struct {
	AcademicProductionListID uint   `json:"academic_production_list_id"`
	AuthorID                 uint   `json:"author_id"`
	Careers                  []uint `json:"careers"`
}

type UpdateAcademicProductionListsAuthorCareersRequest struct {
	AcademicProductionListID uint   `json:"academic_production_list_id"`
	AuthorID                 uint   `json:"author_id"`
	Careers                  []uint `json:"careers"`
}

func (apl *AcademicProductionListsAuthor) TableName() string {
	return model.IndicatorsInformationSchema + ".academic_production_lists_authors"
}

func PostAcademicProductionListsAuthorCareers(request *PostAcademicProductionListsAuthorCareersRequest) (err error) {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		academicProductionListsAuthor := AcademicProductionListsAuthor{
			AcademicProductionListID: request.AcademicProductionListID,
			AuthorID:                 request.AuthorID,
		}
		if err = tx.Create(&academicProductionListsAuthor).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return errors.New("autor ya registrado en la publicación")
			}
			return err
		}

		var academicPeriodID uint
		err = database.DB.Table("indicators_information.academic_production_lists apl").
			Select("apl.academic_period_id").
			Where("apl.id = ?", request.AcademicProductionListID).
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

func GetAcademicProductionListsAuthorsJoinedByAcademicProductionListID(
	academicProductionListID int, response *[]AcademicProductionListsAuthorsCareersJoined) (err error) {
	var academicProductionAuthors []struct {
		Author   string `json:"author"`
		AuthorId uint   `json:"author_id"`
	}
	err = database.DB.Table("indicators_information.academic_production_lists_authors apla").
		Select(`p.name || ' ' || p.lastname as author,
				apla.author_id`).
		Joins("join authors a on apla.author_id = a.id").
		Joins("join people p on a.person_id = p.id").
		Group("apla.author_id, p.name, p.lastname").
		Where("apla.academic_production_list_id = ?", academicProductionListID).
		Scan(&academicProductionAuthors).Error
	if err != nil {
		return err
	}

	for _, authorCareer := range academicProductionAuthors {
		var authorCareers []career.Career
		authorCareers = []career.Career{}
		err = database.DB.Table("indicators_information.academic_production_lists_authors apla").
			Select("c.*").
			Joins("join indicators_information.academic_production_lists apl on apla.academic_production_list_id = apl.id").
			Joins(`join indicators_information.academic_period_author_careers apac
    					on apla.author_id = apac.author_id and apl.academic_period_id = apac.academic_period_id`).
			Joins("join careers c on apac.career_id = c.id").
			Where("apla.author_id = ?", authorCareer.AuthorId).
			Where("apla.academic_production_list_id = ?", academicProductionListID).
			Scan(&authorCareers).Error
		if err != nil {
			return err
		}

		authorCareersJoined := AcademicProductionListsAuthorsCareersJoined{
			AuthorID: authorCareer.AuthorId,
			Author:   authorCareer.Author,
			Careers:  authorCareers,
		}
		*response = append(*response, authorCareersJoined)
	}
	return nil
}

func UpdateAcademicProductionListsAuthorCareersByAcademicPeriodID(
	request *UpdateAcademicProductionListsAuthorCareersRequest) (err error) {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		err = database.DB.Delete(
			&AcademicProductionListsAuthor{},
			"academic_production_list_id = ? and author_id = ?",
			request.AcademicProductionListID, request.AuthorID).
			Error
		if err != nil {
			return err
		}
		//for _, careerID := range request.Careers {
		//	academicProductionListsAuthor := AcademicProductionListsAuthor{
		//		AcademicProductionListID: request.AcademicProductionListID,
		//		AuthorID:                 request.AuthorID,
		//		CareerID:                 careerID,
		//	}
		//	if err := tx.Create(&academicProductionListsAuthor).Error; err != nil {
		//		if errors.Is(err, gorm.ErrForeignKeyViolated) {
		//			return errors.New("claves no encontradas")
		//		}
		//		if errors.Is(err, gorm.ErrDuplicatedKey) {
		//			return errors.New("autor ya registrado")
		//		}
		//		return err
		//	}
		//}
		return nil
	})
}
