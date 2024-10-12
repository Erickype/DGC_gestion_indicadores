package controller

import (
	errorsS "errors"
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/academicPeriod"
	common "github.com/Erickype/DGC_gestion_indicadores_backend/model/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func CreateAcademicPeriod(c *gin.Context) {
	var period model.AcademicPeriod
	err := c.BindJSON(&period)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}
	err = common.CreateAcademicPeriod(&period)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error creando periodo", err)
		return
	}
	c.JSON(http.StatusCreated, period)
}

func GetAcademicPeriods(context *gin.Context) {
	var academicPeriods []model.AcademicPeriod
	err := model.GetAcademicPeriods(&academicPeriods)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando periodos", err)
		return
	}
	context.JSON(http.StatusOK, academicPeriods)
}

func GetAcademicPeriodByID(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}
	var academicPeriod model.AcademicPeriod
	err = model.GetAcademicPeriod(&academicPeriod, int(id))
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando periodo", err)
		return
	}
	context.JSON(http.StatusOK, academicPeriod)
}

func UpdateAcademicPeriod(c *gin.Context) {
	var AcademicPeriod model.AcademicPeriod
	id, _ := strconv.Atoi(c.Param("id"))

	err := model.GetAcademicPeriod(&AcademicPeriod, id)
	if err != nil {
		if errorsS.Is(err, gorm.ErrRecordNotFound) {
			errors.NotFoundResponse(c, "No existe el periodo académico", err)
			return
		}
		errors.InternalServerErrorResponse(c, "Error retornando periodo académico", err)
		return
	}
	err = c.BindJSON(&AcademicPeriod)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}

	err = common.UpdateAcademicPeriod(&AcademicPeriod)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error actualizando periodo", err)
		return
	}
	c.JSON(http.StatusAccepted, AcademicPeriod)
}

func DeleteAcademicPeriod(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}

	err = model.DeleteAcademicPeriod(int(id))
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error eliminando periodo académico", err)
		return
	}
	context.JSON(http.StatusAccepted, errors.CreateCommonDeleteResponse("AcademicPeriod", int(id)))
}
