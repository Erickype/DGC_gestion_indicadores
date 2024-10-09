package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores/DGC_gestion_indicadores_backend/model/academicProduction"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetScienceMagazines(context *gin.Context) {
	var magazines []model.ScienceMagazine
	err := model.GetScienceMagazines(&magazines)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando revistas científicas", err)
		return
	}
	context.JSON(http.StatusOK, magazines)
}
