package controllers

import (
	"net/http"

	"github.com/cyctw/go-crud/models"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

// var userModel = new(models.User)

func (u UserController) Retrieve(c *gin.Context) {

	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "fail"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
