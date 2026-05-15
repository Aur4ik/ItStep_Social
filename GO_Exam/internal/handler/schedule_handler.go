package handler

import (
	"net/http"
	"strconv"

	"project/itStep/internal/models"
	"project/itStep/internal/service"

	"github.com/gin-gonic/gin"
)

func CreateSchedule(c *gin.Context) {

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
		Title       string `json:"title"`
		Description string `json:"description"`
		Date        string `json:"date"`
		Time        string `json:"time"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid data",
		})
		return
	}

	schedule := models.Schedule{
		Title:       input.Title,
		Description: input.Description,
		Date:        input.Date,
		Time:        input.Time,
		TeacherID:   uint(userID),
		CommunityID: uint(communityID),
	}

	err = service.CreateSchedule(&schedule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, schedule)
}

func GetCommunitySchedule(c *gin.Context) {

	idParam := c.Param("id")

	communityID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid community id",
		})
		return
	}

	schedules, err := service.GetCommunitySchedule(
		uint(communityID),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get schedule",
		})
		return
	}

	c.JSON(http.StatusOK, schedules)
}
func UpdateSchedule(c *gin.Context) {

	userID := c.GetInt("user_id")

	role := c.GetString("role")

	idParam := c.Param("id")

	scheduleID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid schedule id",
		})
		return
	}

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Date        string `json:"date"`
		Time        string `json:"time"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid data",
		})
		return
	}

	err = service.UpdateSchedule(
		uint(scheduleID),
		uint(userID),
		role,
		input.Title,
		input.Description,
		input.Date,
		input.Time,
	)

	if err != nil {

		status := http.StatusBadRequest

		if err.Error() == "forbidden" {
			status = http.StatusForbidden
		}

		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "schedule updated",
	})
}

func DeleteSchedule(c *gin.Context) {

	userID := c.GetInt("user_id")

	role := c.GetString("role")

	idParam := c.Param("id")

	scheduleID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid schedule id",
		})
		return
	}

	err = service.DeleteSchedule(
		uint(scheduleID),
		uint(userID),
		role,
	)

	if err != nil {

		status := http.StatusBadRequest

		if err.Error() == "forbidden" {
			status = http.StatusForbidden
		}

		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "schedule deleted",
	})
}