package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/career"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCareer(c *gin.Context) {
	var career model.Career
	err := c.BindJSON(&career)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest, "Error en la solicitud", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	err = model.CreateCareer(&career)
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error creando carrera", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, career)
}

func GetFaculties(context *gin.Context) {
	var careers []model.Career
	err := model.GetCareers(&careers)
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error retornando carreras", err.Error())
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, careers)
}
