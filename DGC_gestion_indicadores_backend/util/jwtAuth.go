package util

import (
	errors "github.com/Erickype/DGC_gestion_indicadores/DGC_gestion_indicadores_backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWTAuth check for valid admin token
func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := ValidateJWT(context)
		if err != nil {
			err := errors.CreateCommonError(http.StatusUnauthorized, "Autenticación requerida", err.Error())
			context.JSON(http.StatusUnauthorized, err)
			context.Abort()
			return
		}
		err = ValidateJWTExpiration(context)
		if err != nil {
			err := errors.CreateCommonError(http.StatusUnauthorized, "Token caducado", err.Error())
			context.JSON(http.StatusUnauthorized, err)
			context.Abort()
			return
		}
		err = ValidateAdminRoleJWT(context)
		if err != nil {
			err := errors.CreateCommonError(http.StatusUnauthorized, "Autenticación admistrador requerida", err.Error())
			context.JSON(http.StatusUnauthorized, err)
			context.Abort()
			return
		}
		context.Next()
	}
}

// JWTAuthUPE check for valid UPE token
func JWTAuthUPE() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := ValidateJWT(context)
		if err != nil {
			err := errors.CreateCommonError(http.StatusUnauthorized, "Autenticación requerida", err.Error())
			context.JSON(http.StatusUnauthorized, err)
			context.Abort()
			return
		}
		err = ValidateJWTExpiration(context)
		if err != nil {
			err := errors.CreateCommonError(http.StatusUnauthorized, "Token caducado", err.Error())
			context.JSON(http.StatusUnauthorized, err)
			context.Abort()
			return
		}
		err = ValidateUPERoleJWT(context)
		if err != nil {
			err := errors.CreateCommonError(http.StatusUnauthorized, "Autenticación UPE requerida", err.Error())
			context.JSON(http.StatusUnauthorized, err)
			context.Abort()
			return
		}
		context.Next()
	}
}
