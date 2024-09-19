package model

import (
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
