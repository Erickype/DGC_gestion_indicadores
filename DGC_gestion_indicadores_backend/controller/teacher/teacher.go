package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/teacher"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateTeacher(c *gin.Context) {
	var Teacher model.Teacher
	err := c.BindJSON(&Teacher)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest, "Error en la solicitud", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	err = model.CreateTeacher(&Teacher)
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error creando docente", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, Teacher)
}

func GetTeachersByAcademicPeriod(context *gin.Context) {
	var Teachers []model.Teacher

	academicPeriodID, err := strconv.ParseInt(context.Param("academicPeriodID"), 10, 64)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest,
			"Error en parámetro academicPeriodID", err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = model.GetTeachersByAcademicPeriod(&Teachers, int(academicPeriodID))
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError,
			"Error retornando docentes en el periodo especificado", err.Error())
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, Teachers)
}
