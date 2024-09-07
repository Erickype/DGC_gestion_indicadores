package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/degree"
	"github.com/gin-gonic/gin"
	"net/http"
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
