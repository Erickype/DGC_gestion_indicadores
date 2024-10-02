package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/booksOrChaptersProduction"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FilterBooksOrChaptersProductionListsByAcademicPeriod(context *gin.Context) {
	var booksOrChaptersProductionListsByAcademicPeriodRequest model.FilterBooksOrChaptersProductionListsByAcademicPeriodRequest
	err := context.BindJSON(&booksOrChaptersProductionListsByAcademicPeriodRequest)
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}
	var booksOrChaptersProductionListsByAcademicPeriodResponse model.FilterBooksOrChaptersProductionListsByAcademicPeriodResponse
	err = model.FilterBooksOrChaptersProductionListsByAcademicPeriod(&booksOrChaptersProductionListsByAcademicPeriodResponse, &booksOrChaptersProductionListsByAcademicPeriodRequest)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando listas de libros o capítulos", err)
		return
	}
	context.JSON(http.StatusOK, booksOrChaptersProductionListsByAcademicPeriodResponse)
}
