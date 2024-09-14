package controller

import (
	errorsS "errors"
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/person"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func CreatePerson(c *gin.Context) {
	var person model.Person
	err := c.BindJSON(&person)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest, "Error en la solicitud", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	err = model.CreatePerson(&person)
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error creando persona", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, person)
}

func GetPersons(context *gin.Context) {
	var persons []model.Person
	err := model.GetPersons(&persons)
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error retornando personas", err.Error())
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, persons)
}

func UpdatePerson(c *gin.Context) {
	var Person model.Person
	id, _ := strconv.Atoi(c.Param("id"))

	err := model.GetPerson(&Person, id)
	if err != nil {
		if errorsS.Is(err, gorm.ErrRecordNotFound) {
			errors.NotFoundResponse(c, "No existe la persona", err)
			return
		}
		errors.InternalServerErrorResponse(c, "Error encontrando persona", err)
		return
	}
	err = c.BindJSON(&Person)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}

	err = model.UpdatePerson(&Person)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error creando persona", err)
		return
	}
	c.JSON(http.StatusAccepted, Person)
}

func DeletePerson(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest,
			"Error en parámetro id.", err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err = model.DeletePerson(int(id))
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError,
			"Error eliminando persona", err.Error())
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "success"})
}
