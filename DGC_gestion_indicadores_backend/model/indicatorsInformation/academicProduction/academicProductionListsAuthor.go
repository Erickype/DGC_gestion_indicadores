package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicProduction "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicProduction"
	"time"
)

type AcademicProductionListsAuthor struct {
	CreatedAt                 time.Time `json:"created_at"`
	UpdatedAt                 time.Time `json:"updated_at"`
	AcademicProductionListDOI string    `gorm:"primary_key" json:"academic_production_list_doi"`
	AuthorID                  uint      `gorm:"primary_key" json:"author_id"`

	AcademicProductionList AcademicProductionList    `json:"-"`
	Author                 academicProduction.Author `json:"-"`
}

func (apl *AcademicProductionListsAuthor) TableName() string {
	return model.IndicatorsInformationSchema + ".academic_production_lists_authors"
}
