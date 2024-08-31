package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/teacher"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FilterTeachers(context *gin.Context) {
	var filter model.FilterTeachersRequest
	err := context.BindJSON(&filter)
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}
	var filterTeachersResponse model.FilterTeachersResponse
	err = model.FilterTeachers(&filterTeachersResponse, &filter)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando profesores", err)
		return
	}
	context.JSON(http.StatusOK, filterTeachersResponse)
}
