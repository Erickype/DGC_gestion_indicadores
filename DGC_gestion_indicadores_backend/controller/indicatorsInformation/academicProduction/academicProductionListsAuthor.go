package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	career "github.com/Erickype/DGC_gestion_indicadores_backend/model/career"
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

func GetAcademicProductionListAuthorPreviousCareers(c *gin.Context) {
	previousCareers := []career.Career{}
	authorID, _ := strconv.Atoi(c.Param("authorID"))
	err := model.GetAcademicProductionListAuthorPreviousCareers(authorID, &previousCareers)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error retornando carreras previas", err)
		return
	}
	c.JSON(http.StatusOK, previousCareers)
}

func GetAcademicProductionListsAuthorsJoinedByAcademicProductionListID(c *gin.Context) {
	academicProductionListID, _ := strconv.Atoi(c.Param("academicProductionListID"))
	response := []model.AcademicProductionListsAuthorsCareersJoined{}
	err := model.GetAcademicProductionListsAuthorsJoinedByAcademicProductionListID(
		academicProductionListID, &response)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error retornando autores", err)
		return
	}
	c.JSON(http.StatusOK, response)
}
