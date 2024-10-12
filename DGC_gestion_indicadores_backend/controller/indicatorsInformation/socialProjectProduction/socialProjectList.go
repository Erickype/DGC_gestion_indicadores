package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/socialProjectProduction"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostSocialProjectList(c *gin.Context) {
	var socialProjectList model.SocialProjectList
	err := c.BindJSON(&socialProjectList)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}
	err = model.PostSocialProjectList(&socialProjectList)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error añadiendo proyecto vinculación a lista", err)
		return
	}
	c.JSON(http.StatusCreated, socialProjectList)
}
