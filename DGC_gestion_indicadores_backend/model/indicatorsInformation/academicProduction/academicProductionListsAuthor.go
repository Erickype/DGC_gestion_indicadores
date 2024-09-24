package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicProduction "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicProduction/author"
	career "github.com/Erickype/DGC_gestion_indicadores_backend/model/career"
	"time"
)

type AcademicProductionListsAuthor struct {
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
	AcademicProductionListID uint      `gorm:"primary_key" json:"academic_production_list_id"`
	AuthorID                 uint      `gorm:"primary_key" json:"author_id"`
	CareerID                 uint      `gorm:"primary_key" json:"career_id"`

	AcademicProductionList AcademicProductionList    `json:"-"`
	Author                 academicProduction.Author `json:"-"`
	Career                 career.Career             `json:"-"`
}

func (apl *AcademicProductionListsAuthor) TableName() string {
	return model.IndicatorsInformationSchema + ".academic_production_lists_authors"
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
