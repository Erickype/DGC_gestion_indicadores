package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/researchInnovationProjectLists"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetResearchInnovationProjectListsByAcademicPeriod(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("academicPeriodID"))
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}
	var researchInnovationProjectLists []model.ResearchInnovationProjectList
	err = model.GetResearchInnovationProjectListByAcademicPeriod(id, &researchInnovationProjectLists)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando listas de proyectos innovación", err)
		return
	}
	if len(researchInnovationProjectLists) == 0 {
		researchInnovationProjectLists = []model.ResearchInnovationProjectList{}
	}
	context.JSON(http.StatusOK, researchInnovationProjectLists)
}
