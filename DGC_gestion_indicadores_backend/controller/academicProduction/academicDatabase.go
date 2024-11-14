package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicProduction"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAcademicDatabases(context *gin.Context) {
	var databases []model.AcademicDatabase
	err := model.GetAcademicDatabases(&databases)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando bases de datos académicas", err)
		return
	}
	context.JSON(http.StatusOK, databases)
}
