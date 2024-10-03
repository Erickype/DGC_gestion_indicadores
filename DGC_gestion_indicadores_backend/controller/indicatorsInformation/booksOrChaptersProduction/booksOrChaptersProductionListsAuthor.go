package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/academicProduction"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostBooksOrChaptersProductionListsAuthorCareers(c *gin.Context) {
	var postAcademicProductionListsAuthorCareersRequest model.PostAcademicProductionListsAuthorCareersRequest
	err := c.BindJSON(&postAcademicProductionListsAuthorCareersRequest)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}
	err = model.PostAcademicProductionListsAuthorCareers(&postAcademicProductionListsAuthorCareersRequest)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error añadiendo autor y sus carreras", err)
		return
	}
	c.JSON(http.StatusCreated, postAcademicProductionListsAuthorCareersRequest)
}
