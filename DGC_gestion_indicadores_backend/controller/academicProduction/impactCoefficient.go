package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicProduction"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetImpactCoefficients(context *gin.Context) {
	var coefficients []model.ImpactCoefficient
	err := model.GetImpactCoefficients(&coefficients)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando coeficientes impacto", err)
		return
	}
	context.JSON(http.StatusOK, coefficients)
}
