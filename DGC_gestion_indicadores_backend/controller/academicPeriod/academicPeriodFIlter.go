package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FilterAcademicPeriods(context *gin.Context) {
	var filterAcademicPeriodsRequest model.FilterAcademicPeriodsRequest
	err := context.BindJSON(&filterAcademicPeriodsRequest)
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}
	var filterAcademicPeriodsResponse model.FilterAcademicPeriodsResponse
	err = model.FilterAcademicPeriods(&filterAcademicPeriodsResponse, &filterAcademicPeriodsRequest)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando periodos académicos", err)
		return
	}
	context.JSON(http.StatusOK, filterAcademicPeriodsResponse)
}
