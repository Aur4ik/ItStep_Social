package handler

import (
	"net/http"
	"strconv"

	"project/itStep/internal/service"

	"github.com/gin-gonic/gin"
)

func AddChatMember(c *gin.Context) {

	chatParam := c.Param("id")

	chatID, err := strconv.Atoi(chatParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid chat id",
		})
		return
	}

	var input struct {
		UserID uint `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid data",
		})
		return
	}

	err = service.AddChatMember(
		uint(chatID),
		input.UserID,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "member added",
	})
}

func GetChatMembers(c *gin.Context) {

	chatParam := c.Param("id")

	chatID, err := strconv.Atoi(chatParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid chat id",
		})
		return
	}

	users, err := service.GetChatMembers(
		uint(chatID),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}