package controller

import (
	errorsS "errors"
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/socialProjectProduction"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

func PutSocialProjectList(c *gin.Context) {
	var socialProjectList model.SocialProjectList
	id, _ := strconv.Atoi(c.Param("id"))

	var socialProjectListGet model.SocialProjectList
	err := model.GetSocialProjectListByID(id, &socialProjectListGet)
	if err != nil {
		if errorsS.Is(err, gorm.ErrRecordNotFound) {
			errors.NotFoundResponse(c, "Proyecto vinculación no encontrado", err)
			return
		}
		errors.InternalServerErrorResponse(c, "Error encontrando proyecto vinculación", err)
		return
	}
	err = c.BindJSON(&socialProjectList)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}

	err = model.PutSocialProjectList(&socialProjectList)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error actualizando proyecto vinculación", err)
		return
	}
	c.JSON(http.StatusAccepted, socialProjectList)
}
