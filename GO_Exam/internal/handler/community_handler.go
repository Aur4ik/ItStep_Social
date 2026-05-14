package handler

import (
	"net/http"
	"strconv"

	"project/itStep/internal/models"
	"project/itStep/internal/service"

	"github.com/gin-gonic/gin"
)

func CreateCommunity(c *gin.Context) {

	userID := c.GetInt("user_id")

	var input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid data",
		})
		return
	}

	community := models.Community{
		Name:        input.Name,
		Description: input.Description,
		OwnerID:     uint(userID),
	}

	err := service.CreateCommunity(&community)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, community)
}

func GetCommunities(c *gin.Context) {

	communities, err := service.GetCommunities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get communities",
		})
		return
	}

	c.JSON(http.StatusOK, communities)
}

func JoinCommunity(c *gin.Context) {

	userID := c.GetInt("user_id")

	idParam := c.Param("id")

	communityID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid community id",
		})
		return
	}

	err = service.JoinCommunity(
		uint(userID),
		uint(communityID),
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "joined community",
	})
}
func GetCommunityPosts(c *gin.Context) {

	idParam := c.Param("id")

	communityID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid community id",
		})
		return
	}

	posts, err := service.GetCommunityPosts(uint(communityID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get posts",
		})
		return
	}

	c.JSON(http.StatusOK, posts)
}
func CreateCommunityPost(c *gin.Context) {

	userID := c.GetInt("user_id")

	idParam := c.Param("id")

	communityID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid community id",
		})
		return
	}

	var input struct {
		Content string `json:"content"`
		Image   string `json:"image"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid data",
		})
		return
	}

	communityIDUint := uint(communityID)

	post := models.Post{
		Content:     input.Content,
		Image:       input.Image,
		AuthorID:    uint(userID),
		CommunityID: &communityIDUint,
	}

	err = service.CreatePost(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, post)
}