package controller

import (
	"net/http"
	"strconv"

	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/evaluationPeriod"
	"github.com/gin-gonic/gin"
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