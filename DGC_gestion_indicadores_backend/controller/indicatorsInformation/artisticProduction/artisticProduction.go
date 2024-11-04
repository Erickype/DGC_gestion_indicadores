package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/artisticProduction"
	"github.com/gin-gonic/gin"
	"net/http"
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
