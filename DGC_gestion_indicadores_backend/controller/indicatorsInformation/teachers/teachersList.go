package controller

import (
	errorsS "errors"
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/teachers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func CreateTeachersList(c *gin.Context) {
	var teachersList model.TeachersList
	err := c.BindJSON(&teachersList)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest, "Error en la solicitud", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	err = model.CreateTeachersList(&teachersList)
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error añadiendo profesor a lista", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, teachersList)
}

func PatchTeachersList(c *gin.Context) {
	var TeachersList model.TeachersList
	academicPeriodID, _ := strconv.Atoi(c.Param("academicPeriodID"))
	teacherID, _ := strconv.Atoi(c.Param("teacherID"))

	err := model.GetTeachersListByAcademicPeriod(academicPeriodID, teacherID)
	if err != nil {
		if errorsS.Is(err, gorm.ErrRecordNotFound) {
			err := errors.CreateCommonError(http.StatusNotFound, "No encontrado", err.Error())
			c.AbortWithStatusJSON(http.StatusNotFound, err)
			return
		}
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error interno", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	err = c.BindJSON(&TeachersList)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest, "Error en la petición", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = model.PatchTeachersLists(&TeachersList)
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error creando persona", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusAccepted, TeachersList)
}
