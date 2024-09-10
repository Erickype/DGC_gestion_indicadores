package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/teachers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PostTeachersListsDegree(c *gin.Context) {
	var teachersListsDegree model.TeachersListsDegree
	err := c.BindJSON(&teachersListsDegree)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}
	err = model.PostTeachersListsDegree(&teachersListsDegree)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error creando título", err)
		return
	}
	c.JSON(http.StatusCreated, teachersListsDegree)
}

func AddDegreeAndTeachersListsDegree(c *gin.Context) {
	var addDegreeAndTeachersListsDegreeRequest model.AddDegreeAndTeachersListsDegreeRequest
	err := c.BindJSON(&addDegreeAndTeachersListsDegreeRequest)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}
	err = model.AddDegreeAndTeachersListsDegree(&addDegreeAndTeachersListsDegreeRequest)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error creando título", err)
		return
	}
	c.JSON(http.StatusCreated, addDegreeAndTeachersListsDegreeRequest)
}

func GetDegreesNotAssigned(context *gin.Context) {
	degreesNotAssignedResponses := []model.GetDegreesNotAssignedResponse{}
	academicPeriodID, _ := strconv.Atoi(context.Param("academicPeriodID"))
	teacherID, _ := strconv.Atoi(context.Param("teacherID"))
	err := model.GetDegreesNotAssigned(&degreesNotAssignedResponses, academicPeriodID, teacherID)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando títulos", err)
		return
	}
	context.JSON(http.StatusOK, degreesNotAssignedResponses)
}

func GetTeachersListsDegreesJoined(context *gin.Context) {
	teachersListsDegreesJoinedResponse := []model.GetTeachersListsDegreesJoinedResponse{}
	academicPeriodID, _ := strconv.Atoi(context.Param("academicPeriodID"))
	teacherID, _ := strconv.Atoi(context.Param("teacherID"))
	err := model.GetTeachersListsDegreesJoined(academicPeriodID, teacherID, &teachersListsDegreesJoinedResponse)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando títulos", err)
		return
	}
	context.JSON(http.StatusOK, teachersListsDegreesJoinedResponse)
}
