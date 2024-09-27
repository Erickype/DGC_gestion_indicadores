package controller

import (
	errorsS "errors"
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicProduction/author"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

func GetAuthorPersonJoinedByAuthorID(context *gin.Context) {
	var authorPerson model.AuthorPerson
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}
	err = model.GetAuthorPersonJoinedByAuthorID(int(id), &authorPerson)
	if err != nil {
		if errorsS.Is(err, gorm.ErrRecordNotFound) {
			errors.NotFoundResponse(context, "Autor no encontrado", err)
			return
		}
		errors.InternalServerErrorResponse(context, "Error retornando author", err)
		return
	}
	context.JSON(http.StatusOK, authorPerson)
}
