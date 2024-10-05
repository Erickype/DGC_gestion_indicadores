package controller

import (
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
