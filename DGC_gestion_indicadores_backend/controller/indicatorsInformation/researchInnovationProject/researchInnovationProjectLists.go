package controller

import (
	errorsS "errors"
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/researchInnovationProjectLists"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

func UpdateResearchInnovationProjectList(c *gin.Context) {
	var researchInnovationProjectList model.ResearchInnovationProjectList
	academicPeriodID, _ := strconv.Atoi(c.Param("academicPeriodID"))

	var researchInnovationProjectListGet model.ResearchInnovationProjectList
	err := model.GetResearchInnovationProjectListByAcademicPeriod(academicPeriodID, &researchInnovationProjectListGet)
	if err != nil {
		if errorsS.Is(err, gorm.ErrRecordNotFound) {
			errors.NotFoundResponse(c, "Proyecto no encontrado", err)
			return
		}
		errors.InternalServerErrorResponse(c, "Error encontrando proyecto innovación", err)
		return
	}
	err = c.BindJSON(&researchInnovationProjectList)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}

	err = model.UpdateResearchInnovationProjectList(&researchInnovationProjectList)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error actualizando proyecto innovación", err)
		return
	}
	c.JSON(http.StatusAccepted, researchInnovationProjectList)
}
