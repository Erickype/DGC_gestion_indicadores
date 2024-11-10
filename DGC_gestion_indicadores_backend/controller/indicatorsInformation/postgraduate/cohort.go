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

func GetPostgraduateCohortListsByProgramID(context *gin.Context) {
	var postgraduateCohortLists []model.PostgraduateCohortList
	programID, err := strconv.Atoi(context.Param("programID"))
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}
	var postgraduateProgram model.PostgraduateProgram
	err = model.GetPostgraduateProgramByID(programID, &postgraduateProgram)
	if err != nil {
		if errorsS.Is(err, gorm.ErrRecordNotFound) {
			errors.NotFoundResponse(context, "Programa posgrado no encontrado", err)
			return
		}
		errors.InternalServerErrorResponse(context, "Error retornando programa posgrado", err)
		return
	}

	err = model.GetPostgraduateCohortListsByProgramID(programID, &postgraduateCohortLists)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando programa posgrado", err)
		return
	}
	if len(postgraduateCohortLists) == 0 {
		postgraduateCohortLists = []model.PostgraduateCohortList{}
	}
	context.JSON(http.StatusOK, postgraduateCohortLists)
}

func GetPostgraduateProgramMissingCohortYearsByProgramID(context *gin.Context) {
	programID, err := strconv.Atoi(context.Param("programID"))
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}
	var postgraduateProgram model.PostgraduateProgram
	err = model.GetPostgraduateProgramByID(programID, &postgraduateProgram)
	if err != nil {
		if errorsS.Is(err, gorm.ErrRecordNotFound) {
			errors.NotFoundResponse(context, "Programa posgrado no encontrado", err)
			return
		}
		errors.InternalServerErrorResponse(context, "Error retornando programa posgrado", err)
		return
	}

	var response model.GetPostgraduateProgramMissingCohortYearsByProgramIDResponse
	err = model.GetPostgraduateProgramMissingCohortYearsByProgramID(programID, &response)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando rango años para program", err)
		return
	}
	if len(response.Years) == 0 {
		response.Years = []int{}
	}
	context.JSON(http.StatusOK, response)
}

func PostPostgraduateCohortList(c *gin.Context) {
	var postgraduateCohortList model.PostgraduateCohortList
	err := c.BindJSON(&postgraduateCohortList)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}
	err = model.PostPostgraduateCohortList(&postgraduateCohortList)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error añadiendo cohorte a lista", err)
		return
	}
	c.JSON(http.StatusCreated, postgraduateCohortList)
}

func UpdatePostgraduateCohortList(c *gin.Context) {
	var postgraduateCohortList model.PostgraduateCohortList
	programID, _ := strconv.Atoi(c.Param("programID"))
	year, _ := strconv.Atoi(c.Param("year"))

	var cohortList model.PostgraduateCohortList
	err := model.GetPostgraduateCohortListByProgramIDAndYear(programID, year, &cohortList)
	if err != nil {
		if errorsS.Is(err, gorm.ErrRecordNotFound) {
			errors.NotFoundResponse(c, "Cohorte no encontrada en programa posgrado", err)
			return
		}
		errors.InternalServerErrorResponse(c, "Error encontrando cohorte en programa posgrado", err)
		return
	}

	err = c.BindJSON(&postgraduateCohortList)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}

	err = model.UpdatePostgraduateCohortList(&postgraduateCohortList)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error actualizando cohorte de programa posgrado", err)
		return
	}
	c.JSON(http.StatusAccepted, postgraduateCohortList)
}
