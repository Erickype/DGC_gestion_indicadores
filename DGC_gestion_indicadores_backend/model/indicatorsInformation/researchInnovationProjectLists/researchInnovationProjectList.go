package model

import (
	"errors"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	academicPeriod "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	"gorm.io/gorm"
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

type ResearchInnovationProjectListJoined struct {
	AcademicPeriodID                 uint   `json:"academic_period_id"`
	AcademicPeriod                   string `json:"academic_period"`
	TotalProjectsUEP                 uint   `json:"total_projects_uep"`
	ProjectsExternalFunding          uint   `json:"projects_external_funding"`
	InternationalCooperationProjects uint   `json:"international_cooperation_projects"`
	NationalCooperationProjects      uint   `json:"national_cooperation_projects"`
}

func (grl ResearchInnovationProjectList) TableName() string {
	return model.IndicatorsInformationSchema + ".research_innovation_project_lists"
}

func GetResearchInnovationProjectLists(
	researchInnovationProjectList *[]ResearchInnovationProjectListJoined) (err error) {
	err = database.DB.Table("indicators_information.research_innovation_project_lists ripl").
		Select(`ripl.academic_period_id,
			ap.name as academic_period,
			ripl.total_projects_uep,
			ripl.projects_external_funding,
			ripl.international_cooperation_projects,
			ripl.national_cooperation_projects`).
		Joins("join academic_periods ap on ripl.academic_period_id = ap.id").
		Order("ripl.updated_at desc").
		Scan(researchInnovationProjectList).Error
	if err != nil {
		return err
	}
	return nil
}

func PostResearchInnovationProjectList(researchInnovationProjectList *ResearchInnovationProjectList) (err error) {
	err = database.DB.Create(&researchInnovationProjectList).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.New("valores para el periodo ya registrados")
		}
		return err
	}
	return nil
}
