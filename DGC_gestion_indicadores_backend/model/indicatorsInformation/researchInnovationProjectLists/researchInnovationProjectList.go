package model

import (
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	"time"
)

type ResearchInnovationProjectList struct {
	CreatedAt                        time.Time `json:"created_at"`
	UpdatedAt                        time.Time `json:"updated_at"`
	AcademicPeriodID                 uint      `gorm:"primary_key" json:"academic_period_id"`
	TotalProjectsUEP                 uint      `json:"total_projects_uep"`
	ProjectsExternalFunding          uint      `json:"projects_external_funding"`
	InternationalCooperationProjects uint      `json:"international_cooperation_projects"`
	NationalCooperationProjects      uint      `json:"national_cooperation_projects"`

	AcademicPeriod academicPeriod.AcademicPeriod `json:"-"`
}

func (grl ResearchInnovationProjectList) TableName() string {
	return model.IndicatorsInformationSchema + ".research_innovation_project_lists"
}

func GetResearchInnovationProjectListByAcademicPeriod(
	academicPeriodID int, researchInnovationProjectList *[]ResearchInnovationProjectList) (err error) {
	err = database.DB.Model(ResearchInnovationProjectList{}).
		Where("academic_period_id = ?", academicPeriodID).
		Order("updated_at desc").
		Scan(researchInnovationProjectList).Error
	if err != nil {
		return err
	}
	return nil
}
