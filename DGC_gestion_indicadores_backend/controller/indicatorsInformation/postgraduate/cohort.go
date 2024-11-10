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
