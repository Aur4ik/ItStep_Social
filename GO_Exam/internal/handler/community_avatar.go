package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"project/itStep/internal/config"

	"github.com/gin-gonic/gin"
)
func CommunityAvatar(c *gin.Context) {

	idParam := c.Param("id")

	communityID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
	})
	return
	}

	file, err := c.FormFile("avatar")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "file not found",
		})
		return
	}

	ext := filepath.Ext(file.Filename)

	allowed := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	if !allowed[ext] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid file type",
		})
		return
}
	if file.Size > 5*1024*1024 {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "file too large",
	})
	return
}

	filename := fmt.Sprintf(
		"%d_%d%s",
		communityID,
		time.Now().Unix(),
		filepath.Ext(file.Filename),
	)

	path := "uploads/avatars/" + filename

	os.MkdirAll("uploads/avatars", os.ModePerm)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to save file",
		})
		return
	}

	avatarURL := "/uploads/avatars/" + filename

	err = config.DB.
		Table("communities").
		Where("id = ?", communityID).
		Update("avatar", avatarURL).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update avatar",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"avatar": avatarURL,
	})
}