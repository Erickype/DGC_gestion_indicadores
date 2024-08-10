package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/person"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FilterPeople(context *gin.Context) {
	var filter model.FilterPeopleRequest
	err := context.BindJSON(&filter)
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}
	var people []model.Person
	err = model.FilterPeople(&people, &filter)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando personas", err)
		return
	}
	context.JSON(http.StatusOK, people)
}
