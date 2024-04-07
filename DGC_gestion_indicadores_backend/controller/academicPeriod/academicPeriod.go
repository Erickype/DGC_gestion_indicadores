package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAcademicPeriod(c *gin.Context) {
	var period model.AcademicPeriod
	err := c.BindJSON(&period)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest, "Error en la solicitud", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	err = model.CreateAcademicPeriod(&period)
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error creando periodo", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, period)
}

func GetAcademicPeriods(context *gin.Context) {
	var academicPeriods []model.AcademicPeriod
	err := model.GetAcademicPeriods(&academicPeriods)
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error retornando periodos", err.Error())
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, academicPeriods)
}
