package controller

import (
	errorsS "errors"
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/artisticProduction"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetArtisticProductionLists(context *gin.Context) {
	var researchInnovationProjectLists []model.ArtisticProductionListJoined
	err := model.GetArtisticProductionLists(&researchInnovationProjectLists)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando listas de producción artística", err)
		return
	}
	if len(researchInnovationProjectLists) == 0 {
		researchInnovationProjectLists = []model.ArtisticProductionListJoined{}
	}
	context.JSON(http.StatusOK, researchInnovationProjectLists)
}

func PostArtisticProductionList(c *gin.Context) {
	var artisticProductionList model.ArtisticProductionList
	err := c.BindJSON(&artisticProductionList)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}
	err = model.PostArtisticProductionList(&artisticProductionList)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error añadiendo valores producción artística a lista", err)
		return
	}
	c.JSON(http.StatusCreated, artisticProductionList)
}

func UpdateArtisticProjectList(c *gin.Context) {
	var artisticProductionList model.ArtisticProductionList
	academicPeriodID, _ := strconv.Atoi(c.Param("academicPeriodID"))

	var artisticProductionListGet model.ArtisticProductionList
	err := model.GetArtisticProductionListByAcademicPeriod(academicPeriodID, &artisticProductionListGet)
	if err != nil {
		if errorsS.Is(err, gorm.ErrRecordNotFound) {
			errors.NotFoundResponse(c, "Producción artística no encontrada", err)
			return
		}
		errors.InternalServerErrorResponse(c, "Error encontrando producción artística", err)
		return
	}
	err = c.BindJSON(&artisticProductionList)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}

	err = model.UpdateArtisticProductionList(&artisticProductionList)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error actualizando producción artística", err)
		return
	}
	c.JSON(http.StatusAccepted, artisticProductionList)
}
