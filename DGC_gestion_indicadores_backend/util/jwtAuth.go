package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JWTAuth check for valid admin token
func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := ValidateJWT(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}
		err = ValidateAdminRoleJWT(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Only Administrator is allowed to perform this action"})
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
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			context.Abort()
			return
		}
		err = ValidateUPERoleJWT(context)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Only registered Customers are allowed to perform this action"})
			context.Abort()
			return
		}
		context.Next()
	}
}
