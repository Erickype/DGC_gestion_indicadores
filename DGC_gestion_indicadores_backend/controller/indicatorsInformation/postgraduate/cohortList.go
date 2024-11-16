package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/postgraduate"
	"github.com/gin-gonic/gin"
	"net/http"
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
