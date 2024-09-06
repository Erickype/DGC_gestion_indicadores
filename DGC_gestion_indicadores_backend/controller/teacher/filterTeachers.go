package controller

import (
	errorsS "errors"
	errors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/teacher"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func FilterTeachers(context *gin.Context) {
	var filter model.FilterTeachersRequest
	err := context.BindJSON(&filter)
	if err != nil {
		errors.BadRequestResponse(context, err)
		return
	}
	var filterTeachersResponse model.FilterTeachersResponse
	err = model.FilterTeachers(&filterTeachersResponse, &filter)
	if err != nil {
		errors.InternalServerErrorResponse(context, "Error retornando profesores", err)
		return
	}
	context.JSON(http.StatusOK, filterTeachersResponse)
}

func GetTeacherPersonJoinedByTeacherID(context *gin.Context) {
	var teacherPersonJoined model.TeacherPerson
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		err := errors.CreateCommonError(http.StatusBadRequest,
			"Error en parámetro teacher id.", err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	err = model.GetTeacherPersonJoinedByTeacherID(int(id), &teacherPersonJoined)
	if err != nil {
		if errorsS.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, teacherPersonJoined)
}
