package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/researchInnovationProjectLists"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetResearchInnovationProjectListsByAcademicPeriod(context *gin.Context) {
	var researchInnovationProjectLists []model.ResearchInnovationProjectListJoined
	err := model.GetResearchInnovationProjectLists(&researchInnovationProjectLists)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando listas de proyectos innovación", err)
		return
	}
	if len(researchInnovationProjectLists) == 0 {
		researchInnovationProjectLists = []model.ResearchInnovationProjectListJoined{}
	}
	context.JSON(http.StatusOK, researchInnovationProjectLists)
}
