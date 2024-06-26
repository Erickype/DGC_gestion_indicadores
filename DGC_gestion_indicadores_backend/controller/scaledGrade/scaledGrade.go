package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/scaledGrade"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateScaledGrade(c *gin.Context) {
	var ScaledGrade model.ScaledGrade
	err := c.BindJSON(&ScaledGrade)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest, "Error en la solicitud", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	err = model.CreateScaledGrade(&ScaledGrade)
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error creando grado escalafonado.", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, ScaledGrade)
}

func GetScaledGrades(context *gin.Context) {
	var ScaledGrades []model.ScaledGrade
	err := model.GetScaledGrades(&ScaledGrades)
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error retornando grado escalafonado.", err.Error())
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, ScaledGrades)
}
