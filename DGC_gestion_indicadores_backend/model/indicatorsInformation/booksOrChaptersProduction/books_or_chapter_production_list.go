package model

import (
	"errors"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	knowledgeField "github.com/Erickype/DGC_gestion_indicadores_backend/model/knowledgeField"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type BooksOrChaptersProductionList struct {
	ID               uint           `gorm:"primaryKey"`
	DOI              string         `gorm:"unique" json:"DOI"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	IsChapter        *bool          `json:"is_chapter"`
	Title            string         `gorm:"size:1000;not null;" json:"title"`
	PublicationDate  datatypes.Date `gorm:"not null;" json:"publication_date"`
	PeerReviewed     *bool          `json:"peer_reviewed"`
	AcademicPeriodID uint           `json:"academic_period_id"`
	DetailedFieldID  uint           `json:"detailed_field_id"`

	DetailedField  knowledgeField.DetailedField  `json:"-"`
	AcademicPeriod academicPeriod.AcademicPeriod `json:"-"`
}

func (bc BooksOrChaptersProductionList) TableName() string {
	return model.IndicatorsInformationSchema + ".books_or_chapter_production_lists"
}

func GetBooksOrChaptersProductionListByID(id int, response *BooksOrChaptersProductionList) (err error) {
	err = database.DB.First(&response, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("no existe el libro o capítulo en el periodo")
		}
		return err
	}
	return nil
}

func PostBooksOrChaptersProductionList(request *BooksOrChaptersProductionList) (err error) {
	err = database.DB.Create(request).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("libro o capítulo en lista ya existe")
		}
		return err
	}
	return nil
}

func UpdateBooksOrChaptersProductionList(booksOrChaptersProductionList *BooksOrChaptersProductionList) (err error) {
	err = database.DB.Model(&BooksOrChaptersProductionList{}).
		Where("id = ?",
			booksOrChaptersProductionList.ID).
		Updates(booksOrChaptersProductionList).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("libro o capítulo en lista ya existe")
		}
		return err
	}
	return nil
}
