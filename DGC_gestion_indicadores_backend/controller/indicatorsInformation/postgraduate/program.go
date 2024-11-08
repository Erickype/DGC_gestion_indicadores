package controller

import (
	errorsS "errors"
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/postgraduate"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

func UpdatePostGraduateProgram(c *gin.Context) {
	var postgraduateProgram model.PostgraduateProgram
	programID, _ := strconv.Atoi(c.Param("programID"))

	var checkResponse model.PostgraduateProgram
	err := model.GetPostgraduateProgramByID(programID, &checkResponse)
	if err != nil {
		if errorsS.Is(err, gorm.ErrRecordNotFound) {
			errors.NotFoundResponse(c, "Programa posgrado no encontrado", err)
			return
		}
		errors.InternalServerErrorResponse(c, "Error encontrando programa posgrado", err)
		return
	}
	err = c.BindJSON(&postgraduateProgram)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}

	err = model.UpdatePostgraduateProgram(&postgraduateProgram)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error actualizando programa posgrado", err)
		return
	}
	c.JSON(http.StatusAccepted, postgraduateProgram)
}
