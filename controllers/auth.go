package controllers

import (
	"fmt"
	"net/http"

	"github.com/cyctw/go-crud/models"
	"github.com/gin-gonic/gin"
)

type UserInput struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type AuthController struct{}

func (auth AuthController) Signup(c *gin.Context) {
	// Signup
	var userInput UserInput
	if err := c.ShouldBind(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	fmt.Println("acount", userInput)

	user := models.User{Name: userInput.Username, Password: userInput.Password}
	fmt.Println("user", user)

	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "GG"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (auth AuthController) Login(c *gin.Context) {
	// Login
	var userInput UserInput
	if err := c.ShouldBind(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bind fail"})
		return
	}
	fmt.Println(userInput)
	var user models.User
	if err := models.DB.Where("name = ? AND password = ?", userInput.Username, userInput.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "can't find"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user.ID})

}

func (auth AuthController) Logout(c *gin.Context) {
	// Logout
}
