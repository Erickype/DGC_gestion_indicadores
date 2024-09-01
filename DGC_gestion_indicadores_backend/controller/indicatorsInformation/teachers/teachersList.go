package teachers

import (
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/indicatorsInformation/teachers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTeachersList(c *gin.Context) {
	var teachersList model.TeachersList
	err := c.BindJSON(&teachersList)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest, "Error en la solicitud", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	err = model.CreateTeachersList(&teachersList)
	if err != nil {
		err := errors.CreateCommonError(http.StatusInternalServerError, "Error añadiendo profesor a lista", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, teachersList)
}
