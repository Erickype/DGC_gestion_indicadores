package controller

import (
	errorsS "errors"
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/knowledgeField"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

func GetDetailedFilterJoinedByDetailedFieldID(context *gin.Context) {
	var detailedFieldJoined model.DetailedFieldJoined
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}
	err = model.GetDetailedFilterJoinedByDetailedFieldID(int(id), &detailedFieldJoined)
	if err != nil {
		if errorsS.Is(err, gorm.ErrRecordNotFound) {
			errors.NotFoundResponse(context, "Campo detallado no encontrado", err)
			return
		}
		errors.InternalServerErrorResponse(context, "Error retornando campos detallados", err)
		return
	}
	context.JSON(http.StatusOK, detailedFieldJoined)
}
