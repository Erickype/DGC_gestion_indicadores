package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/academicProduction"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FilterAcademicProductionListsByAcademicPeriod(context *gin.Context) {
	var filter model.FilterAcademicProductionListsByAcademicPeriodRequest
	err := context.BindJSON(&filter)
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}
	var filterAcademicProductionListsByAcademicPeriodResponse model.FilterAcademicProductionListsByAcademicPeriodResponse
	err = model.FilterAcademicProductionListsByAcademicPeriod(&filterAcademicProductionListsByAcademicPeriodResponse, &filter)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando listas de publicaciones académicas", err)
		return
	}
	context.JSON(http.StatusOK, filterAcademicProductionListsByAcademicPeriodResponse)
}
