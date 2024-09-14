package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicProduction/author"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FilterAuthors(context *gin.Context) {
	var filter model.FilterAuthorsRequest
	err := context.BindJSON(&filter)
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}
	var filterAuthorsResponse model.FilterAuthorsResponse
	err = model.FilterAuthors(&filterAuthorsResponse, &filter)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando autores", err)
		return
	}
	context.JSON(http.StatusOK, filterAuthorsResponse)
}
