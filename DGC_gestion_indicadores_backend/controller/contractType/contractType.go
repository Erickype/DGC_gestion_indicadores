package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/contractType"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetContractTypes(context *gin.Context) {
	var ContractTypes []model.ContractType
	err := model.GetContractTypes(&ContractTypes)
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error retornando tipos contrato.", err.Error())
		context.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, ContractTypes)
}
