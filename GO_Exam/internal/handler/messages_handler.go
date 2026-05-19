package handler

import (
	"net/http"
	"strconv"

	"project/itStep/internal/models"
	"project/itStep/internal/service"

	"github.com/gin-gonic/gin"
)

func CreateMessage(c *gin.Context) {
	userID := c.GetInt("user_id")

	chatID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid chat id"})
		return
	}

	var input struct {
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	message := models.Message{
		Content:  input.Content,
		ChatID:   uint(chatID),
		SenderID: uint(userID),
	}

	if err := service.CreateMessage(&message); err != nil {

		if err.Error() == "not a chat member" {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, message)
}

func GetChatMessages(c *gin.Context) {
	chatID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid chat id"})
		return
	}

	messages, err := service.GetChatMessages(uint(chatID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get messages"})
		return
	}

	c.JSON(http.StatusOK, messages)
}
