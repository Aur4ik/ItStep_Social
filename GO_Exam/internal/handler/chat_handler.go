package handler

import (
	"net/http"
	"strconv"

	"project/itStep/internal/models"
	"project/itStep/internal/service"

	"github.com/gin-gonic/gin"
)

func CreateChat(c *gin.Context) {

	var input struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid data",
		})
		return
	}

	chat := models.Chat{
		Name: input.Name,
	}

	err := service.CreateChat(&chat)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, chat)
}

func GetChats(c *gin.Context) {

	chats, err := service.GetChats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get chats",
		})
		return
	}

	c.JSON(http.StatusOK, chats)
}

func UpdateChat(c *gin.Context) {

	idParam := c.Param("id")

	chatID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid chat id",
		})
		return
	}

	var input struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid data",
		})
		return
	}

	err = service.UpdateChat(
		uint(chatID),
		input.Name,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "chat updated",
	})
}

func DeleteChat(c *gin.Context) {

	idParam := c.Param("id")

	chatID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid chat id",
		})
		return
	}

	err = service.DeleteChat(uint(chatID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "chat deleted",
	})
}