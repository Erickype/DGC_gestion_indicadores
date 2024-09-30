package controller

import (
	errorsS "errors"
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/academicProduction"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

func PatchAcademicProductionList(c *gin.Context) {
	var academicProductionList model.AcademicProductionList
	id, _ := strconv.Atoi(c.Param("id"))

	var checkResponse model.AcademicProductionList
	err := model.GetAcademicProductionListByID(id, &checkResponse)
	if err != nil {
		if errorsS.Is(err, gorm.ErrRecordNotFound) {
			errors.NotFoundResponse(c, "Publicación no encontrada", err)
			return
		}
		errors.InternalServerErrorResponse(c, "Error encontrando publicación académica", err)
		return
	}
	err = c.BindJSON(&academicProductionList)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}

	err = model.PatchAcademicProductionList(&academicProductionList)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error actualizando publicación académica", err)
		return
	}
	c.JSON(http.StatusAccepted, academicProductionList)
}
