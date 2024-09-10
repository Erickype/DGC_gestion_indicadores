package model

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommonError struct {
	StatusCode int    `json:"status_code,omitempty"`
	Status     string `json:"status,omitempty"`
	Message    string `json:"message,omitempty"`
	Detail     string `json:"detail,omitempty"`
}

func CreateCommonError(statusCode int, message string, detail string) *CommonError {
	commonError := &CommonError{
		StatusCode: statusCode,
		Status:     http.StatusText(statusCode),
		Message:    message,
		Detail:     detail,
	}
	return commonError
}

func BadRequestResponse(context *gin.Context, err error) {
	commonError := CreateCommonError(http.StatusBadRequest, "Error en la petición", err.Error())
	context.AbortWithStatusJSON(http.StatusBadRequest, commonError)
	return
}

func InternalServerErrorResponse(context *gin.Context, message string, err error) {
	commonError := CreateCommonError(http.StatusInternalServerError, message, err.Error())
	context.AbortWithStatusJSON(http.StatusInternalServerError, commonError)
	return
}

func NotFoundResponse(context *gin.Context, message string, err error) {
	commonError := CreateCommonError(http.StatusNotFound, message, err.Error())
	context.AbortWithStatusJSON(http.StatusNotFound, commonError)
	return
}
