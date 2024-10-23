package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicProduction/author"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostAuthor(context *gin.Context) {
	var author model.Author
	err := context.BindJSON(&author)
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}
	err = model.PostAuthor(&author)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error añadiendo autor", err)
		return
	}
	context.JSON(http.StatusCreated, author)
}
