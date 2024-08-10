package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	academicProduction "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicProduction"
	"gorm.io/datatypes"
	"time"
)

type BooksOrChaptersProductionList struct {
	DOI               uint           `gorm:"primary_key"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	PublicationTypeID uint           `json:"publication_type_id"`
	AcademicPeriodID  uint           `json:"academic_period_id"`
	BookTitle         string         `gorm:"size:1000;not null;" json:"book_title"`
	ChapterTitle      string         `gorm:"size:1000;not null;" json:"chapter_title"`
	PublicationDate   datatypes.Date `gorm:"not null;" json:"publication_date"`
	PeerReviewed      bool           `json:"peer_reviewed"`

	PublicationType academicProduction.PublicationType `json:"-"`
	AcademicPeriod  academicPeriod.AcademicPeriod      `json:"-"`
}

func (bc BooksOrChaptersProductionList) TableName() string {
	return model.IndicatorsInformationSchema + ".books_or_chapter_production_lists"
}
