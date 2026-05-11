package handler

import (
	"net/http"

	"project/itStep/internal/models"
	"project/itStep/internal/service"
	"project/itStep/internal/utils"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid data",
		})
		return
	}

	err := service.Register(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.Password = ""

	c.JSON(http.StatusCreated, user)
}

func Login(c *gin.Context) {

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid data",
		})
		return
	}

	user, err := service.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(int(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "token error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func Me(c *gin.Context) {

	userID := c.GetInt("user_id")

	user, err := service.GetUserByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	user.Password = ""

	c.JSON(http.StatusOK, user)
}