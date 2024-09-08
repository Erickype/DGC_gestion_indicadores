package controller

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/teachers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddDegreeAndTeachersListsDegree(c *gin.Context) {
	var addDegreeAndTeachersListsDegreeRequest model.AddDegreeAndTeachersListsDegreeRequest
	err := c.BindJSON(&addDegreeAndTeachersListsDegreeRequest)
	if err != nil {
		errors.BadRequestResponse(c, err)
		return
	}
	err = model.AddDegreeAndTeachersListsDegree(&addDegreeAndTeachersListsDegreeRequest)
	if err != nil {
		errors.InternalServerErrorResponse(c, "Error creando título", err)
		return
	}
	c.JSON(http.StatusCreated, addDegreeAndTeachersListsDegreeRequest)
}
