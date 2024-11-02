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

func PostResearchInnovationProjectList(c *gin.Context) {
	var researchInnovationProjectList model.ResearchInnovationProjectList
	err := c.BindJSON(&researchInnovationProjectList)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}
	err = model.PostResearchInnovationProjectList(&researchInnovationProjectList)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error añadiendo valores proyectos innovación a lista", err)
		return
	}
	c.JSON(http.StatusCreated, researchInnovationProjectList)
}
