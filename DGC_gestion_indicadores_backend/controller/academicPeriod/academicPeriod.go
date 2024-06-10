package controller

import (
	errorsS "errors"
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	common "github.com/Erickype/DGC_gestion_indicadores_backend/model/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func CreateAcademicPeriod(c *gin.Context) {
	var period model.AcademicPeriod
	err := c.BindJSON(&period)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest, "Error en la solicitud", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	err = common.CreateAcademicPeriod(&period)
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

func UpdateAcademicPeriod(c *gin.Context) {
	var AcademicPeriod model.AcademicPeriod
	id, _ := strconv.Atoi(c.Param("id"))

	err := model.GetAcademicPeriod(&AcademicPeriod, id)
	if err != nil {
		if errorsS.Is(err, gorm.ErrRecordNotFound) {
			err := errors.CreateCommonError(http.StatusNotFound, "No existe el periodo académico", err.Error())
			c.AbortWithStatusJSON(http.StatusNotFound, err)
			return
		}
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error interno", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	err = c.BindJSON(&AcademicPeriod)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest, "Error en la petición", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = common.UpdateAcademicPeriod(&AcademicPeriod)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func DeleteAcademicPeriod(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest,
			"Error en parámetro id.", err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = model.DeleteAcademicPeriod(int(id))
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError,
			"Error eliminando periodo evaluación", err.Error())
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "success"})
}
