package controller

import (
	"errors"
	model "github.com/Erickype/DGC_gestion_indicadores_backend/model/auth"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateRole creates a new role
func CreateRole(c *gin.Context) {
	var Role model.Role
	err := c.BindJSON(&Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
	}
	err = model.CreateRole(&Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Role)
}

// GetRoles returns the list of roles
func GetRoles(c *gin.Context) {
	var Role []model.Role
	err := model.GetRoles(&Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Role)
}

// GetRole return a role based on the param "id" from route
func GetRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var Role model.Role
	err := model.GetRole(&Role, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Role)
}

// UpdateRole updates a role based on the param "id" from route and the body
func UpdateRole(c *gin.Context) {
	var Role model.Role
	id, _ := strconv.Atoi(c.Param("id"))
	err := model.GetRole(&Role, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = c.BindJSON(&Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
	}
	err = model.UpdateRole(&Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Role)
}
