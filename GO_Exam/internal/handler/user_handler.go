package handler

import (
	"net/http"

	"strconv"

	"project/itStep/internal/service"

	"github.com/gin-gonic/gin"
)

func UpdateUserRole(c *gin.Context){
	idParam := c.Param("id")

	userID, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid user id",
		})
		return
	}
	var input struct{
		Role string `json: "role"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid data",
		})
		return
	}

	err = service.UpdateUserRole(uint(userID), input.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : "role updated",
	})
}