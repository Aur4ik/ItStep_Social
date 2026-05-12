package handler

import (
	"net/http"
	"strconv"

	"project/itStep/internal/repository"
	"project/itStep/internal/service"

	"github.com/gin-gonic/gin"
)

func ToggleLike(c *gin.Context){
	userID := c.GetInt("user_id")

	postIDParam := c.Param("id")

	postId, err := strconv.Atoi(postIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid post id",
		})
		return
	}
	message, err := service.ToggleLike(
		uint(userID),
		uint(postId),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "Failed to toggle like",
		})
		return
	}

	count, err := repository.CountLikes(uint(postId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "failed to count likes",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : message,
		"like_count" : count,
	})
}

