package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	career "github.com/Erickype/DGC_gestion_indicadores_backend/model/career"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAcademicPeriodAuthorPreviousCareers(c *gin.Context) {
	var previousCareers []career.Career
	authorID, _ := strconv.Atoi(c.Param("authorID"))
	err := model.GetAcademicPeriodAuthorPreviousCareers(authorID, &previousCareers)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error retornando carreras previas", err)
		return
	}
	if len(previousCareers) == 0 {
		previousCareers = []career.Career{}
	}
	c.JSON(http.StatusOK, previousCareers)
}
