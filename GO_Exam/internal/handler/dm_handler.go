package handler

import (
	"net/http"

	"project/itStep/internal/service"

	"github.com/gin-gonic/gin"
)

func CreateDMChat(c *gin.Context) {

	currentUser := c.GetInt("user_id")

	var input struct {
		UserID uint `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid data",
		})
		return
	}

	chat, err := service.CreateDMChat(
		uint(currentUser),
		input.UserID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create dm",
		})
		return
	}

	c.JSON(http.StatusOK, chat)
}