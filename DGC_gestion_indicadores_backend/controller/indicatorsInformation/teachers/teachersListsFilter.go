package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/teachers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FilterTeachersLists(context *gin.Context) {
	var filter model.FilterTeachersListsByAcademicPeriodRequest
	err := context.BindJSON(&filter)
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}
	var filterTeachersResponse model.FilterTeachersListsByAcademicPeriodResponse
	err = model.FilterTeachersListsByAcademicPeriod(&filterTeachersResponse, &filter)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando listas de profesores", err)
		return
	}
	context.JSON(http.StatusOK, filterTeachersResponse)
}
