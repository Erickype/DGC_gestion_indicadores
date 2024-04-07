package model

import "net/http"

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
