package handler

import (
	"net/http"
	"strconv"

	"project/itStep/internal/models"
	"project/itStep/internal/repository"
	"project/itStep/internal/service"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	userID := c.GetInt("user_id")

	var input struct {
		Content string `json:"content"`
		Image   string `json:"image"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	post := models.Post{
		Content:  input.Content,
		Image:    input.Image,
		AuthorID: uint(userID),
	}

	if err := service.CreatePost(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdPost, err := repository.GetPostWithAuthor(post.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load post"})
		return
	}

	c.JSON(http.StatusCreated, createdPost)
}

func GetPosts(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	offset := (page - 1) * limit

	posts, err := service.GetPosts(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func DeletePost(c *gin.Context) {
	userID := c.GetInt("user_id")

	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = service.DeletePost(uint(postID), uint(userID))
	if err != nil {

		if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post deleted"})
}

func GetPostByID(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	post, err := service.GetPostByID(uint(postID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}
