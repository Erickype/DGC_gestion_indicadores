package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	author "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicProduction/author"
	career "github.com/Erickype/DGC_gestion_indicadores_backend/model/career"
)

type AcademicPeriodAuthorCareer struct {
	AcademicPeriodID uint `gorm:"primary_key" json:"academic_period_id"`
	AuthorID         uint `gorm:"primary_key" json:"author_id"`
	CareerID         uint `gorm:"primary_key" json:"career_id"`

	AcademicPeriod academicPeriod.AcademicPeriod `json:"-"`
	Author         author.Author                 `json:"-"`
	Career         career.Career                 `json:"-"`
}

func (apac AcademicPeriodAuthorCareer) TableName() string {
	return model.IndicatorsInformationSchema + ".academic_period_author_careers"
}

func GetAcademicPeriodAuthorPreviousCareers(authorID int, previousCareers *[]career.Career) (err error) {
	var academicPeriodsIDs []uint
	database.DB.Table("academic_periods ap").
		Order("ap.start_date desc").
		Select("ap.id").
		Limit(1).Scan(&academicPeriodsIDs)
	if len(academicPeriodsIDs) == 0 {
		database.DB.Table("academic_periods ap").
			Order("ap.start_date desc").
			Select("ap.id").
			Limit(2).Scan(&academicPeriodsIDs)
	}

	var authorPreviousCareers []uint
	database.DB.Table("indicators_information.academic_period_author_careers apac").
		Distinct("apac.career_id").
		Where("apac.academic_period_id in (?)", academicPeriodsIDs).
		Where("apac.author_id = ?", authorID).Scan(&authorPreviousCareers)

	err = database.DB.Table("careers c").
		Where("c.id in (?)", authorPreviousCareers).
		Scan(&previousCareers).Error

	if err != nil {
		return err
	}
	return nil
}
