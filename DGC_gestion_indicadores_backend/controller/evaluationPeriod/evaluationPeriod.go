package controller

import (
	errorsS "errors"
	"net/http"
	"strconv"

	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/evaluationPeriod"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateEvaluationPeriod(c *gin.Context) {
	var period model.EvaluationPeriod
	err := c.BindJSON(&period)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest, "Error en la solicitud", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = model.CreateEvaluationPeriod(&period)
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error creando periodo", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, period)
}

func GetEvaluationPeriods(context *gin.Context) {
	var EvaluationPeriods []model.EvaluationPeriod
	err := model.GetEvaluationPeriods(&EvaluationPeriods)
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error retornando periodos", err.Error())
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, EvaluationPeriods)
}

func UpdateEvaluationPeriod(c *gin.Context) {
	var EvaluationPeriod model.EvaluationPeriod
	id, _ := strconv.Atoi(c.Param("id"))

	err := model.GetEvaluationPeriod(&EvaluationPeriod, id)
	if err != nil {
		if errorsS.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = c.BindJSON(&EvaluationPeriod)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = model.UpdateEvaluationPeriod(&EvaluationPeriod)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func DeleteEvaluationPeriod(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest,
			"Error en parámetro id.", err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = model.DeleteEvaluationPeriod(int(id))
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError,
			"Error eliminando periodo evaluación", err.Error())
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "success"})
}