package handler

import (
	"net/http"
	"strconv"

	"project/itStep/internal/service"

	"github.com/gin-gonic/gin"
)

func ToggleLike(c *gin.Context) {
	userID := c.GetInt("user_id")

	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}


	message, count, err := service.ToggleLike(uint(userID), uint(postID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to toggle like"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    message,
		"like_count": count,
	})
}
