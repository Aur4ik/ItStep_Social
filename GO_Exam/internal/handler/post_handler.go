package handler

import (
	"net/http"
	"strconv"

	"project/itStep/internal/models"
	"project/itStep/internal/service"
	"project/itStep/internal/repository"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context){
	userID := c.GetInt("user_id")

	var input struct {
		Content string `json:"content"`
		Image   string `json:"image"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid data",
		})
		return
	}
	post := models.Post{
		Content: input.Content,
		Image: input.Image,
		AuthorID: uint(userID),
	}

	err := service.CreatePost(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
	createdPost, err := repository.GetPostWithAuthor(post.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to load author",
		})
		return
	}

	c.JSON(http.StatusCreated, createdPost)
}

func GetPosts(c *gin.Context){

	posts, err := service.GetPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : "failed to get posts",
		})
		return
	}
	c.JSON(http.StatusOK, posts)
}
func DeletePost(c *gin.Context) {

	userID := c.GetInt("user_id")

	idParam := c.Param("id")

	postID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	err = service.DeletePost(uint(postID), uint(userID))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "post deleted",
	})
}