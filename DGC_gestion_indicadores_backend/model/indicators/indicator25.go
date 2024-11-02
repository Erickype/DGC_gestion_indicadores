package model

import (
	"errors"
	"fmt"
	"github.com/Erickype/DGC_gestion_indicadores_backend/database"
	"github.com/Erickype/DGC_gestion_indicadores_backend/model"
	researchInnovationProject "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/researchInnovationProjectLists"
)

func CalculateIndicator25(evaluationPeriodID int, indicator *IndicatorsEvaluationPeriod) (err error) {
	var academicPeriodIds []int
	err = getAcademicPeriodsInEvaluationPeriod(evaluationPeriodID, &academicPeriodIds)
	if err != nil {
		return err
	}

	var valuesSums researchInnovationProject.ResearchInnovationProjectList
	err = database.DB.Table("indicators_information.research_innovation_project_lists ripl").
		Select(`    sum(ripl.total_projects_uep) as total_projects_uep,
			sum(ripl.projects_external_funding) as projects_external_funding,
			sum(ripl.international_cooperation_projects) as international_cooperation_projects,
			sum(ripl.national_cooperation_projects) as national_cooperation_projects`).
		Where("ripl.academic_period_id in ?", academicPeriodIds).
		Scan(&valuesSums).Error
	if err != nil {
		return err
	}
	if valuesSums.TotalProjectsUEP <= 0 {
		return errors.New(fmt.Sprintf("indicador %d: no hay total proyectos UEP", 25))
	}

	indicator.ActualValue = (float64(valuesSums.ProjectsExternalFunding) +
		float64(valuesSums.InternationalCooperationProjects) +
		float64(valuesSums.NationalCooperationProjects)) /
		float64(valuesSums.TotalProjectsUEP) * 100

	indicator.TargetValue = model.Indicator25TargetValue
	err = RefreshIndicatorEvaluationPeriod(indicator)
	if err != nil {
		return err
	}
	return nil
}
