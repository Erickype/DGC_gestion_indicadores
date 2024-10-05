package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/academicProduction"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PostAcademicProductionListsAuthorCareers(c *gin.Context) {
	var postAcademicProductionListsAuthorCareersRequest model.PostAcademicProductionListsAuthorCareersRequest
	err := c.BindJSON(&postAcademicProductionListsAuthorCareersRequest)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}
	err = model.PostAcademicProductionListsAuthorCareers(&postAcademicProductionListsAuthorCareersRequest)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error añadiendo autor y sus carreras", err)
		return
	}
	c.JSON(http.StatusCreated, postAcademicProductionListsAuthorCareersRequest)
}

func UpdateAcademicProductionListsAuthorCareersByAcademicPeriodID(c *gin.Context) {
	var request model.UpdateAcademicProductionListsAuthorCareersRequest
	err := c.BindJSON(&request)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}
	err = model.UpdateAcademicProductionListsAuthorCareersByAcademicPeriodID(&request)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error carreras del autor", err)
		return
	}
	c.JSON(http.StatusAccepted, request)
}

func GetAcademicProductionListsAuthorsJoinedByAcademicProductionListID(c *gin.Context) {
	academicProductionListID, _ := strconv.Atoi(c.Param("academicProductionListID"))
	var response []model.AcademicProductionListsAuthorsCareersJoined
	err := model.GetAcademicProductionListsAuthorsJoinedByAcademicProductionListID(
		academicProductionListID, &response)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error retornando autores", err)
		return
	}
	if len(response) == 0 {
		response = []model.AcademicProductionListsAuthorsCareersJoined{}
	}
	c.JSON(http.StatusOK, response)
}
