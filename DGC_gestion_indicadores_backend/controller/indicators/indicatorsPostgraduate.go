package controller

import (
	errors2 "errors"
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

func GetCalculateIndicatorsByPostgraduateCohortYear(context *gin.Context) {
	var response []model.IndicatorsPostgraduateJoined
	cohortYear, _ := strconv.Atoi(context.Param("cohortYear"))
	errs := model.GetCalculateByPostgraduateCohortYear(cohortYear, &response)
	if len(errs) > 0 {
		var resultError string
		for _, err := range errs {
			resultError += err.Error() + "\n"
		}
		errors.InternalServerErrorResponse(context, "Error calculando indicadores", errors2.New(resultError))
		return
	}
	if len(response) == 0 {
		response = []model.IndicatorsPostgraduateJoined{}
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
