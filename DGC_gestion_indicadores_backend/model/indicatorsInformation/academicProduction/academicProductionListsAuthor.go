package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicProduction "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicProduction/author"
	career "github.com/Erickype/DGC_gestion_indicadores_backend/model/career"
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
		/*for _, careerID := range request.Careers {
			academicProductionListsAuthor := AcademicProductionListsAuthor{
				AcademicProductionListID: request.AcademicProductionListID,
				AuthorID:                 request.AuthorID,
				CareerID:                 careerID,
			}
			if err := tx.Create(&academicProductionListsAuthor).Error; err != nil {
				if errors.Is(err, gorm.ErrForeignKeyViolated) {
					return errors.New("claves no encontradas")
				}
				if errors.Is(err, gorm.ErrDuplicatedKey) {
					return errors.New("autor ya registrado")
				}
				return err
			}
		}*/
		return nil
	})
}

func GetAcademicProductionListAuthorPreviousCareers(authorID int, previousCareers *[]career.Career) (err error) {
	var academicPeriodsIDs []uint
	database.DB.Table("academic_periods ap").
		Order("ap.start_date desc").
		Select("ap.id").
		Limit(2).Scan(&academicPeriodsIDs)

	var authorPreviousCareers []uint
	database.DB.Table("indicators_information.academic_production_lists_authors apla").
		Distinct("apla.career_id").
		Joins("join indicators_information.academic_production_lists apl on apla.academic_production_list_id = apl.id").
		Where("apl.academic_period_id in (?)", academicPeriodsIDs).
		Where("apla.author_id = ?", authorID).Scan(&authorPreviousCareers)

	err = database.DB.Table("careers c").
		Where("c.id in (?)", authorPreviousCareers).
		Scan(&previousCareers).Error

	if err != nil {
		return err
	}
	return nil
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
		authorCareers := []career.Career{}
		err = database.DB.Table("indicators_information.academic_production_lists_authors apla").
			Select("c.*").
			Joins("join careers c on apla.career_id = c.id").
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
