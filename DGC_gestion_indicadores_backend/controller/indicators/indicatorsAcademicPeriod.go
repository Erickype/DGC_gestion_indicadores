package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicators"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCalculateIndicatorByTypeID(context *gin.Context) {
	var response model.IndicatorsAcademicPeriod
	academicPeriodID, _ := strconv.Atoi(context.Param("academicPeriodID"))
	indicatorTypeID, _ := strconv.Atoi(context.Param("indicatorTypeID"))
	err := model.CalculateIndicatorByTypeIDAndAcademicPeriod(academicPeriodID, indicatorTypeID, &response)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error calculando indicador", err)
		return
	}
	context.JSON(http.StatusOK, response)
}

func GetIndicatorsByAcademicPeriod(context *gin.Context) {
	var response []model.IndicatorsAcademicPeriod
	academicPeriodID, _ := strconv.Atoi(context.Param("academicPeriodID"))
	err := model.GetIndicatorsByAcademicPeriod(academicPeriodID, &response)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error obteniendo indicador", err)
		return
	}
	context.JSON(http.StatusOK, response)
}
