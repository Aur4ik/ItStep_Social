package handler

import (
	"net/http"
	"strconv"

	"project/itStep/internal/models"
	"project/itStep/internal/service"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context){
	userID := c.GetInt("user_id")

	postIDParam := c.Param("id")

	postID, err := strconv.Atoi(postIDParam)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid post id",
		})
		return 
	}
	var input struct {
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid data",
		})
		return
	}

	comment := models.Comment{
		Content: input.Content,
		AuthorID: uint(userID),
		PostID: uint(postID),
	}

	err = service.CreateComment(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})

		return
	}
	c.JSON(http.StatusCreated, comment)
}

func GetComments(c *gin.Context){
	postIDParam := c.Param("id")

	postID, err := strconv.Atoi(postIDParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid post id",
		})
		return
	}

	comments, err := service.GetCommentByPostID(uint(postID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Feiled to get comments",
		})
		return
	}

	c.JSON(http.StatusOK, comments)
}