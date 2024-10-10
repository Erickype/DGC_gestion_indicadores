package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	"time"
)

type SocialProjectList struct {
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	AcademicPeriodID uint      `gorm:"primary_key" json:"academic_period_id"`
	CareerID         uint      `gorm:"primary_key" json:"career_id"`
	Name             string    `gorm:"not null;unique;size:1000" json:"name"`
}

func (sp SocialProjectList) TableName() string {
	return model.IndicatorsInformationSchema + ".social_project_lists"
}
