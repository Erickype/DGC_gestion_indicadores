package controller

import (
	errorsS "errors"
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/postgraduate"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func PostFilterCohortLists(context *gin.Context) {
	var filterCohortListsRequest model.FilterCohortListsRequest
	err := context.BindJSON(&filterCohortListsRequest)
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}
	var filterCohortListsResponse model.FilterCohortListsResponse
	err = model.PostFilterCohortLists(&filterCohortListsResponse, &filterCohortListsRequest)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando año de cohorte", err)
		return
	}
	context.JSON(http.StatusOK, filterCohortListsResponse)
}

func GetCohortListByYear(context *gin.Context) {
	var cohortList model.CohortList
	year, err := strconv.Atoi(context.Param("year"))
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}
	err = model.GetCohortListByYear(year, &cohortList)
	if err != nil {
		if errorsS.Is(err, gorm.ErrRecordNotFound) {
			errors.NotFoundResponse(context, "Cohorte no encontrada", err)
			return
		}
		errors.InternalServerErrorResponse(context, "Error retornando cohorte", err)
		return
	}
	context.JSON(http.StatusOK, cohortList)
}
