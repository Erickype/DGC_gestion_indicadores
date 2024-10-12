package model

import (
	"errors"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	career "github.com/Erickype/DGC_gestion_indicadores_backend/model/career"
	"gorm.io/gorm"
	"time"
)

type SocialProjectList struct {
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	ID               uint      `gorm:"primary_key"`
	AcademicPeriodID uint      `json:"academic_period_id"`
	CareerID         uint      `json:"career_id"`
	Name             string    `gorm:"not null;unique;size:1000" json:"name"`

	AcademicPeriod academicPeriod.AcademicPeriod `json:"-"`
	Career         career.Career                 `json:"-"`
}

func (sp SocialProjectList) TableName() string {
	return model.IndicatorsInformationSchema + ".social_project_lists"
}

func PostSocialProjectList(socialProjectList *SocialProjectList) (err error) {
	err = database.DB.Create(socialProjectList).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("proyecto vinculación en lista ya existe")
		}
		return err
	}
	return nil
}
