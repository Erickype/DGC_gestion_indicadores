package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/faculty"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateFaculty(c *gin.Context) {
	var period model.Faculty
	err := c.BindJSON(&period)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest, "Error en la solicitud", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	err = model.CreateFaculty(&period)
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error creando facultad", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, period)
}

func GetFaculties(context *gin.Context) {
	var faculties []model.Faculty
	err := model.GetFaculties(&faculties)
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error retornando periodos", err.Error())
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, faculties)
}
