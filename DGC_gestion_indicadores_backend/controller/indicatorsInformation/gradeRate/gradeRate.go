package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/gradeRate"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetGradeRateListsByAcademicPeriod(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		errors.BadRequestResponse(context, err)
	}
	var gradeRateListsJoined []model.GradeRateListJoined
	err = model.GetGradeRateListsByAcademicPeriod(id, &gradeRateListsJoined)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando listas de tasas de grado", err)
		return
	}
	if len(gradeRateListsJoined) == 0 {
		gradeRateListsJoined = []model.GradeRateListJoined{}
	}
	context.JSON(http.StatusOK, gradeRateListsJoined)
}

func PostGradeRateList(c *gin.Context) {
	var gradeRateList model.GradeRateList
	err := c.BindJSON(&gradeRateList)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}
	err = model.PostGradeRateList(&gradeRateList)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error añadiendo valores de tasa de grado a lista", err)
		return
	}
	c.JSON(http.StatusCreated, gradeRateList)
}
