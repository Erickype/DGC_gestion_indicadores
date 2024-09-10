package controller

import (
	errorsS "errors"
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/degree"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func PostTeachersDegree(c *gin.Context) {
	var teachersDegree model.TeachersDegree
	err := c.BindJSON(&teachersDegree)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}
	err = model.PostTeachersDegree(&teachersDegree)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error creando título", err)
		return
	}
	c.JSON(http.StatusCreated, teachersDegree)
}

func PatchTeachersDegreeByTeachersDegreeID(context *gin.Context) {
	var response model.TeachersDegree
	id, _ := strconv.Atoi(context.Param("id"))
	err := model.GetTeachersDegreeByID(&response, id)
	if err != nil {
		if errorsS.Is(err, gorm.ErrRecordNotFound) {
			errors.NotFoundResponse(context, "Título no encontrado", err)
			return
		}
		errors.InternalServerErrorResponse(context, "Error encontrando título", err)
		return
	}

	var request model.PatchTeachersDegreeByTeachersDegreeIDRequest
	err = context.BindJSON(&request)
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}
	err = model.PatchTeachersDegreeByTeachersDegreeID(&request, id)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error actualizando título", err)
		return
	}
	context.JSON(http.StatusAccepted, response)
}
