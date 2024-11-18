package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicators"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCalculateIndicatorByTypeIDAndCohortYear(context *gin.Context) {
	var response model.IndicatorsPostgraduate
	cohortYear, _ := strconv.Atoi(context.Param("cohortYear"))
	indicatorTypeID, _ := strconv.Atoi(context.Param("indicatorTypeID"))
	err := model.CalculateIndicatorByTypeIDAndCohortYear(cohortYear, indicatorTypeID, &response)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error calculando indicador", err)
		return
	}
	context.JSON(http.StatusOK, response)
}

func GetIndicatorsByPostgraduateCohortYear(context *gin.Context) {
	var response []model.IndicatorsPostgraduateJoined
	cohortYear, _ := strconv.Atoi(context.Param("cohortYear"))
	err := model.GetIndicatorsByPostgraduateCohortYear(cohortYear, &response)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error obteniendo indicadores", err)
		return
	}
	if len(response) == 0 {
		response = []model.IndicatorsPostgraduateJoined{}
	}
	context.JSON(http.StatusOK, response)
}
