package controller

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	knowledgeField "github.com/Erickype/DGC_gestion_indicadores_backend/model/knowledgeField"
	"gorm.io/datatypes"
	"time"
)

type BooksOrChaptersProductionList struct {
	ID               uint           `gorm:"primaryKey"`
	DOI              string         `gorm:"unique" json:"DOI"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	IsChapter        bool           `json:"is_chapter"`
	Title            string         `gorm:"size:1000;not null;" json:"title"`
	PublicationDate  datatypes.Date `gorm:"not null;" json:"publication_date"`
	PeerReviewed     bool           `json:"peer_reviewed"`
	AcademicPeriodID uint           `json:"academic_period_id"`
	DetailedFieldID  uint           `json:"detailed_field_id"`

	DetailedField  knowledgeField.DetailedField  `json:"-"`
	AcademicPeriod academicPeriod.AcademicPeriod `json:"-"`
}

func (bc BooksOrChaptersProductionList) TableName() string {
	return model.IndicatorsInformationSchema + ".books_or_chapter_production_lists"
}
