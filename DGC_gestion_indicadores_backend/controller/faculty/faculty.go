package controller

import (
	errorsS "errors"
	common "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/faculty"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func CreateFaculty(c *gin.Context) {
	var faculty model.Faculty
	err := c.BindJSON(&faculty)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest, "Error en la solicitud", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	err = model.CreateFaculty(&faculty)
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error creando facultad", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, faculty)
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

func UpdateFaculty(context *gin.Context) {
	var faculty model.Faculty
	id, _ := strconv.Atoi(context.Param("id"))
	err := model.GetFaculty(&faculty, id)
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
	err = context.BindJSON(&faculty)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest, "Error en la petición", err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = model.UpdateFaculty(&faculty)
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error creando facultad", err.Error())
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusAccepted, faculty)
}

func DeleteFaculty(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest,
			"Error en la solicitud", err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = model.DeleteFaculty(int(id))
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error eliminando facultad", err.Error())
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusAccepted, common.CreateCommonDeleteResponse("Faculty", int(id)))
}
