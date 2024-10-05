package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/booksOrChaptersProduction"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func PostBooksOrChaptersProductionListsAuthorCareers(c *gin.Context) {
	var postBooksOrChaptersProductionListsAuthorCareersRequest model.PostBooksOrChaptersProductionListsAuthorCareersRequest
	err := c.BindJSON(&postBooksOrChaptersProductionListsAuthorCareersRequest)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}
	err = model.PostBooksOrChaptersProductionListsAuthorCareers(&postBooksOrChaptersProductionListsAuthorCareersRequest)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error añadiendo autor y sus carreras", err)
		return
	}
	c.JSON(http.StatusCreated, postBooksOrChaptersProductionListsAuthorCareersRequest)
}

func UpdateBooksOrChaptersProductionListsAuthorCareers(c *gin.Context) {
	var request model.UpdateBooksOrChaptersProductionListsAuthorCareersRequest
	err := c.BindJSON(&request)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}
	err = model.UpdateBooksOrChaptersProductionListsAuthorCareers(&request)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error actualizando carreras del autor", err)
		return
	}
	c.JSON(http.StatusAccepted, request)
}

func GetBooksOrChaptersProductionListsAuthorsJoinedByBooksOrChaptersProductionListID(c *gin.Context) {
	bookOrChaptersProductionListID, _ := strconv.Atoi(c.Param("booksOrChaptersProductionListID"))
	var response []model.BooksOrChaptersProductionListsAuthorsCareersJoined
	err := model.GetBooksOrChaptersProductionListsAuthorsJoinedByBooksOrChaptersProductionListID(
		bookOrChaptersProductionListID, &response)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error retornando autores", err)
		return
	}
	if len(response) == 0 {
		response = []model.BooksOrChaptersProductionListsAuthorsCareersJoined{}
	}
	c.JSON(http.StatusOK, response)
}
