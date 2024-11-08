package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/postgraduate"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FilterPostGraduatePrograms(context *gin.Context) {
	var filterPostgraduateProgramsRequest model.FilterPostgraduateProgramsRequest
	err := context.BindJSON(&filterPostgraduateProgramsRequest)
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}
	var filterPostGraduateProgramsResponse model.FilterPostGraduateProgramsResponse
	err = model.FilterPostGraduatePrograms(&filterPostGraduateProgramsResponse, &filterPostgraduateProgramsRequest)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando programas posgrado", err)
		return
	}
	context.JSON(http.StatusOK, filterPostGraduateProgramsResponse)
}

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
