package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/postgraduate"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostPostgraduateProgram(c *gin.Context) {
	var postgraduateProgram model.PostgraduateProgram
	err := c.BindJSON(&postgraduateProgram)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}
	err = model.PostPostgraduateProgram(&postgraduateProgram)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error añadiendo programa posgrado", err)
		return
	}
	c.JSON(http.StatusCreated, postgraduateProgram)
}
