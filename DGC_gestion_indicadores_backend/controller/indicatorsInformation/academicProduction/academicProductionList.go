package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/academicProduction"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostAcademicProductionList(c *gin.Context) {
	var academicProductionList model.AcademicProductionList
	err := c.BindJSON(&academicProductionList)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}
	err = model.PostAcademicProductionList(&academicProductionList)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error añadiendo publicación a lista", err)
		return
	}
	c.JSON(http.StatusCreated, academicProductionList)
}
