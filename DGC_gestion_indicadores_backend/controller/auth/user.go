package controller

import (
	"errors"
	"fmt"
	commonErrors "github.com/Erickype/DGC_gestion_indicadores_backend/model"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/auth"
	"github.com/Erickype/DGC_gestion_indicadores_backend/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// Register user
func Register(context *gin.Context) {
	var input model.Register

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
		RoleID:   3,
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})

}

// Login creates the session for user
func Login(context *gin.Context) {
	var input model.Login

	if err := context.ShouldBindJSON(&input); err != nil {
		var errorMessage string
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			validationError := validationErrors[0]
			if validationError.Tag() == "required" {
				errorMessage = fmt.Sprintf("%s not provided", validationError.Field())
			}
		}
		err := commonErrors.CreateCommonError(http.StatusBadRequest, "Error en la solicitud", errorMessage)
		context.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := model.GetUserByUsername(input.Username)

	if err != nil {
		err := commonErrors.CreateCommonError(http.StatusBadRequest, "Error en la solicitud", err.Error())
		context.JSON(http.StatusBadRequest, err)
		return
	}

	err = user.ValidateUserPassword(input.Password)

	if err != nil {
		err := commonErrors.CreateCommonError(http.StatusBadRequest, "Error ingresando", "credenciales incorrectas")
		context.JSON(http.StatusBadRequest, err)
		return
	}

	jwt, err := util.GenerateJWT(user)
	if err != nil {
		err := commonErrors.CreateCommonError(http.StatusBadRequest, "Error interno", "error generando token")
		context.JSON(http.StatusBadRequest, err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"token":    jwt,
		"username": input.Username,
		"message":  "Successfully logged in"})

}

// GetUsers returns a list of the users
func GetUsers(context *gin.Context) {
	var user []model.User
	err := model.GetUsers(&user)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, user)
}

// GetUser return a user by the route param "id"
func GetUser(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	var user model.User
	err := model.GetUser(&user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithStatus(http.StatusNotFound)
			return
		}

		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	context.JSON(http.StatusOK, user)
}

// UpdateUser a user info based on route param "id" and body
func UpdateUser(c *gin.Context) {
	var Update model.Update
	var User model.User
	id, _ := strconv.Atoi(c.Param("id"))

	err := model.GetUser(&User, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			commonErrors.NotFoundResponse(c, "Usuario no encontrado", err)
			return
		}
		commonErrors.InternalServerErrorResponse(c, "Error encontrando usuario", err)
		return
	}
	err = c.BindJSON(&Update)
	if err != nil {
		commonErrors.BadRequestResponse(c, err)
		return
	}
	User.Username = Update.Username
	User.Email = Update.Email
	User.RoleID = Update.RoleID

	err = model.UpdateUser(&User)
	if err != nil {
		commonErrors.InternalServerErrorResponse(c, "Error actualizando usuario", err)
		return
	}
	c.JSON(http.StatusAccepted, User)
}
