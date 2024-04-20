package controller

import (
	errorsS "errors"
	"net/http"
	"strconv"

	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/teacher"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	var Teachers []model.GetJoin = []model.GetJoin{}

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

func UpdateTeacher(c *gin.Context) {
	var Update model.Update
	var Teacher model.Teacher
	id, _ := strconv.Atoi(c.Param("id"))

	err := model.GetTeacher(&Teacher, id)
	if err != nil {
		if errorsS.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = c.BindJSON(&Update)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Teacher.ID = Update.ID
	Teacher.AcademicPeriodID = Update.AcademicPeriodID
	Teacher.PersonID = Update.PersonID
	Teacher.CareerID = Update.CareerID
	Teacher.DedicationID = Update.DedicationID
	Teacher.ScaledGradeID = Update.ScaledGradeID

	err = model.UpdateTeacher(&Teacher)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func DeleteTeacher(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest,
			"Error en parámetro teacher id.", err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = model.DeleteTeacher(int(id))
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError,
			"Error eliminando docente", err.Error())
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "success"})
}
