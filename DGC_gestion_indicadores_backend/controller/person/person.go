package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/person"
	"github.com/gin-gonic/gin"
	"net/http"
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
