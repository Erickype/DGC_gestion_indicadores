package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/socialProjectProduction"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FilterSocialProjectListsByAcademicPeriod(context *gin.Context) {
	var filterSocialProjectListsByAcademicPeriodRequest model.FilterSocialProjectListsByAcademicPeriodRequest
	err := context.BindJSON(&filterSocialProjectListsByAcademicPeriodRequest)
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}
	var filterSocialProjectListsByAcademicPeriodResponse model.FilterSocialProjectListsByAcademicPeriodResponse
	err = model.FilterSocialProjectListsByAcademicPeriod(&filterSocialProjectListsByAcademicPeriodResponse, &filterSocialProjectListsByAcademicPeriodRequest)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando listas de proyectos vinculación", err)
		return
	}
	context.JSON(http.StatusOK, filterSocialProjectListsByAcademicPeriodResponse)
}
