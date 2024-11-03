package controller

import (
	errors2 "errors"
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicators"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCalculateIndicatorByTypeIDAndEvaluationPeriod(context *gin.Context) {
	var response model.IndicatorsEvaluationPeriod
	evaluationPeriodID, _ := strconv.Atoi(context.Param("evaluationPeriodID"))
	indicatorTypeID, _ := strconv.Atoi(context.Param("indicatorTypeID"))
	err := model.CalculateIndicatorByTypeIDAndEvaluationPeriod(evaluationPeriodID, indicatorTypeID, &response)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error calculando indicador", err)
		return
	}
	context.JSON(http.StatusOK, response)
}

func GetIndicatorByTypeIDAndEvaluationPeriodJoined(context *gin.Context) {
	var response model.IndicatorEvaluationPeriodJoined
	evaluationPeriodID, _ := strconv.Atoi(context.Param("evaluationPeriodID"))
	indicatorTypeID, _ := strconv.Atoi(context.Param("indicatorTypeID"))
	err := model.GetIndicatorByTypeIDAndEvaluationPeriodJoined(evaluationPeriodID, indicatorTypeID, &response)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando indicador", err)
		return
	}
	context.JSON(http.StatusOK, response)
}

func GetCalculateIndicatorsByEvaluationPeriod(context *gin.Context) {
	var response []model.IndicatorEvaluationPeriodJoined
	evaluationPeriodID, _ := strconv.Atoi(context.Param("evaluationPeriodID"))
	errs := model.CalculateIndicatorsByEvaluationPeriod(evaluationPeriodID, &response)
	if len(errs) > 0 {
		var resultError string
		for _, err := range errs {
			resultError += err.Error() + "\n"
		}
		errors.InternalServerErrorResponse(context, "Error calculando indicadores", errors2.New(resultError))
		return
	}
	if len(response) == 0 {
		response = []model.IndicatorEvaluationPeriodJoined{}
	}
	context.JSON(http.StatusOK, response)
}

func GetIndicatorsByEvaluationPeriod(context *gin.Context) {
	var response []model.IndicatorEvaluationPeriodJoined
	evaluationPeriodID, _ := strconv.Atoi(context.Param("evaluationPeriodID"))
	err := model.GetIndicatorsByEvaluationPeriod(evaluationPeriodID, &response)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error obteniendo indicadores", err)
		return
	}
	if len(response) == 0 {
		response = []model.IndicatorEvaluationPeriodJoined{}
	}
	context.JSON(http.StatusOK, response)
}
