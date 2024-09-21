package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/knowledgeField"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FilterDetailedFields(context *gin.Context) {
	var filter model.FilterDetailedFieldRequest
	err := context.BindJSON(&filter)
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}
	var filterDetailedFieldResponse model.FilterDetailedFieldResponse
	err = model.FilterDetailedFields(&filterDetailedFieldResponse, &filter)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando campos detallados", err)
		return
	}
	context.JSON(http.StatusOK, filterDetailedFieldResponse)
}
