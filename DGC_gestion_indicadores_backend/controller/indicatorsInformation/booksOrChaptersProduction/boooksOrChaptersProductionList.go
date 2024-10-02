package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/booksOrChaptersProduction"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostBooksOrChaptersProductionList(c *gin.Context) {
	var booksOrChaptersProductionList model.BooksOrChaptersProductionList
	err := c.BindJSON(&booksOrChaptersProductionList)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}
	err = model.PostBooksOrChaptersProductionList(&booksOrChaptersProductionList)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error añadiendo libro o capítulo a lista", err)
		return
	}
	c.JSON(http.StatusCreated, booksOrChaptersProductionList)
}
