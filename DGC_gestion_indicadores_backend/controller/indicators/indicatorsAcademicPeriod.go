package controller

import (
	errors2 "errors"
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicators"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCalculateIndicatorByTypeIDAndAcademicPeriod(context *gin.Context) {
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

func GetCalculateIndicatorsByAcademicPeriod(context *gin.Context) {
	var response []model.IndicatorAcademicPeriodJoined
	academicPeriodID, _ := strconv.Atoi(context.Param("academicPeriodID"))
	errs := model.CalculateIndicatorsByAcademicPeriod(academicPeriodID, &response)
	if len(errs) > 0 {
		var resultError string
		for _, err := range errs {
			resultError += err.Error() + "\n"
		}
		errors.InternalServerErrorResponse(context, "Error calculando indicadores", errors2.New(resultError))
		return
	}
	if len(response) == 0 {
		response = []model.IndicatorAcademicPeriodJoined{}
	}
	context.JSON(http.StatusOK, response)
}

func GetIndicatorsByAcademicPeriod(context *gin.Context) {
	var response []model.IndicatorAcademicPeriodJoined
	academicPeriodID, _ := strconv.Atoi(context.Param("academicPeriodID"))
	err := model.GetIndicatorsByAcademicPeriod(academicPeriodID, &response)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error obteniendo indicador", err)
		return
	}
	if len(response) == 0 {
		response = []model.IndicatorAcademicPeriodJoined{}
	}
	context.JSON(http.StatusOK, response)
}
