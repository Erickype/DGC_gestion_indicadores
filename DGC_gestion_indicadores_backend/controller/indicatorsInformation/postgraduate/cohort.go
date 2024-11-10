package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/postgraduate"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostPostgraduateCohortList(c *gin.Context) {
	var postgraduateCohortList model.PostgraduateCohortList
	err := c.BindJSON(&postgraduateCohortList)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}
	err = model.PostPostgraduateCohortList(&postgraduateCohortList)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error añadiendo cohorte a lista", err)
		return
	}
	c.JSON(http.StatusCreated, postgraduateCohortList)
}
