package controller

import (
	errorsS "errors"
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/career"
	facultyModel "github.com/Erickype/DGC_gestion_indicadores_backend/model/faculty"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func CreateCareer(c *gin.Context) {
	var career model.Career
	err := c.BindJSON(&career)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest, "Error en la solicitud", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)

		return
	}
	var faculty facultyModel.Faculty
	err = facultyModel.GetFaculty(&faculty, int(career.FacultyID))
	if err != nil {
		err := errors.CreateCommonError(http.StatusNotFound, "Error en solicitud", err.Error())
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}
	err = model.CreateCareer(&career)
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error creando carrera", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, career)
}

func GetCareers(context *gin.Context) {
	var careers []model.Career
	err := model.GetCareers(&careers)
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error retornando carreras", err.Error())
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, careers)
}

func UpdateCareer(context *gin.Context) {
	var career model.Career
	id, _ := strconv.Atoi(context.Param("id"))
	err := model.GetCareer(&career, id)
	if err != nil {
		if errorsS.Is(err, gorm.ErrRecordNotFound) {
			err := errors.CreateCommonError(http.StatusNotFound, "No existe la facultad", err.Error())
			context.AbortWithStatusJSON(http.StatusNotFound, err)
			return
		}
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error interno", err.Error())
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	err = context.BindJSON(&career)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest, "Error en la petición", err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = model.UpdateCareer(&career)
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error creando facultad", err.Error())
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusAccepted, career)
}
