package controller

import (
	errorsS "errors"
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/booksOrChaptersProduction"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

func UpdateBooksOrChaptersProductionList(c *gin.Context) {
	var booksOrChaptersProductionList model.BooksOrChaptersProductionList
	id, _ := strconv.Atoi(c.Param("id"))

	var checkResponse model.BooksOrChaptersProductionList
	err := model.GetBooksOrChaptersProductionListByID(id, &checkResponse)
	if err != nil {
		if errorsS.Is(err, gorm.ErrRecordNotFound) {
			errors.NotFoundResponse(c, "Libro o capítulo no encontrado", err)
			return
		}
		errors.InternalServerErrorResponse(c, "Error encontrando libro o capítulo", err)
		return
	}
	err = c.BindJSON(&booksOrChaptersProductionList)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}

	err = model.UpdateBooksOrChaptersProductionList(&booksOrChaptersProductionList)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error actualizando libro o capítulo", err)
		return
	}
	c.JSON(http.StatusAccepted, booksOrChaptersProductionList)
}
