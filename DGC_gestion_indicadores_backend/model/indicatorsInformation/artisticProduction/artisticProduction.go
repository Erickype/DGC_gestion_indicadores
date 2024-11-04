package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	"time"
)

type ArtisticProductionList struct {
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt                 time.Time `json:"updated_at"`
	AcademicPeriodID          uint      `gorm:"primary_key" json:"academic_period_id"`
	InternationalArtisticWork uint      `json:"international_artistic_work"`
	NationalArtisticWork      uint      `json:"national_artistic_work"`

	AcademicPeriod academicPeriod.AcademicPeriod `json:"-"`
}

type ArtisticProductionListJoined struct {
	AcademicPeriodID          uint   `json:"academic_period_id"`
	AcademicPeriod            string `json:"academic_period"`
	InternationalArtisticWork uint   `json:"international_artistic_work"`
	NationalArtisticWork      uint   `json:"national_artistic_work"`
}

func (grl ArtisticProductionList) TableName() string {
	return model.IndicatorsInformationSchema + ".artistic_production_lists"
}

func GetArtisticProductionLists(
	researchInnovationProjectList *[]ArtisticProductionListJoined) (err error) {
	err = database.DB.Table("indicators_information.artistic_production_lists apl").
		Select(`apl.academic_period_id,
			ap.name as academic_period,
			apl.international_artistic_work,
			apl.national_artistic_work`).
		Joins("join academic_periods ap on apl.academic_period_id = ap.id").
		Order("apl.updated_at desc").
		Scan(researchInnovationProjectList).Error
	if err != nil {
		return err
	}
	return nil
}
