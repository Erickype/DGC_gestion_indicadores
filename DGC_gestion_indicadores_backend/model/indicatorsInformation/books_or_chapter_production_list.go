package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	knowledgeField "github.com/Erickype/DGC_gestion_indicadores_backend/model/knowledgeField"
	"gorm.io/datatypes"
	"time"
)

type BooksOrChaptersProductionList struct {
	DOI              uint           `gorm:"primary_key"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DetailedFieldID  uint           `json:"detailed_field_id"`
	AcademicPeriodID uint           `json:"academic_period_id"`
	BookTitle        string         `gorm:"size:1000;not null;" json:"book_title"`
	ChapterTitle     string         `gorm:"size:1000;not null;" json:"chapter_title"`
	PublicationDate  datatypes.Date `gorm:"not null;" json:"publication_date"`
	PeerReviewed     bool           `json:"peer_reviewed"`

	DetailedField  knowledgeField.DetailedField  `json:"-"`
	AcademicPeriod academicPeriod.AcademicPeriod `json:"-"`
}

func (bc BooksOrChaptersProductionList) TableName() string {
	return model.IndicatorsInformationSchema + ".books_or_chapter_production_lists"
}
