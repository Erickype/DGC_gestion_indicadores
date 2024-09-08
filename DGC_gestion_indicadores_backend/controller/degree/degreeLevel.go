package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/degree"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDegreeLevels(context *gin.Context) {
	var degreeLevels []model.DegreeLevel
	err := model.GetDegreeLevels(&degreeLevels)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando niveles", err)
		return
	}
	context.JSON(http.StatusOK, degreeLevels)
}
