package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	career "github.com/Erickype/DGC_gestion_indicadores_backend/model/career"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/academicProduction"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAcademicProductionListAuthorPreviousCareers(c *gin.Context) {
	previousCareers := []career.Career{}
	authorID, _ := strconv.Atoi(c.Param("authorID"))
	err := model.GetAcademicProductionListAuthorPreviousCareers(authorID, &previousCareers)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error retornando carreras previas", err)
		return
	}
	c.JSON(http.StatusOK, previousCareers)
}
