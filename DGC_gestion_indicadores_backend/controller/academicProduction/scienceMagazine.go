package controller

import (
	errorsS "errors"
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicProduction"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func PostFilterScienceMagazines(context *gin.Context) {
	var filter model.FilterScienceMagazinesRequest
	err := context.BindJSON(&filter)
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}
	var filterScienceMagazinesResponse model.FilterScienceMagazinesResponse
	err = model.PostFilterScienceMagazines(&filterScienceMagazinesResponse, &filter)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando revistas científicas", err)
		return
	}
	context.JSON(http.StatusOK, filterScienceMagazinesResponse)
}

func PostScienceMagazine(c *gin.Context) {
	var scienceMagazine model.ScienceMagazine
	err := c.BindJSON(&scienceMagazine)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}
	err = model.PostScienceMagazine(&scienceMagazine)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error añadiendo revista científica", err)
		return
	}
	c.JSON(http.StatusCreated, scienceMagazine)
}

func UpdateScienceMagazine(c *gin.Context) {
	var scienceMagazine model.ScienceMagazine
	id, _ := strconv.Atoi(c.Param("id"))

	var checkResponse model.ScienceMagazine
	err := model.GetScienceMagazineByID(id, &checkResponse)
	if err != nil {
		if errorsS.Is(err, gorm.ErrRecordNotFound) {
			errors.NotFoundResponse(c, "Revista científica no encontrada", err)
			return
		}
		errors.InternalServerErrorResponse(c, "Error encontrando revista científica ", err)
		return
	}
	err = c.BindJSON(&scienceMagazine)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}

	err = model.PutScienceMagazine(&scienceMagazine)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error actualizando revista científica", err)
		return
	}
	c.JSON(http.StatusAccepted, scienceMagazine)
}

func GetScienceMagazineFilterJoinedByScienceMagazineID(context *gin.Context) {
	var magazines model.ScienceMagazinesJoined
	id, _ := strconv.Atoi(context.Param("id"))
	err := model.GetScienceMagazineFilterJoinedByScienceMagazineID(id, &magazines)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando revista científica", err)
		return
	}
	context.JSON(http.StatusOK, magazines)
}

func GetScienceMagazines(context *gin.Context) {
	var magazines []model.ScienceMagazine
	err := model.GetScienceMagazines(&magazines)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando revistas científicas", err)
		return
	}
	context.JSON(http.StatusOK, magazines)
}
