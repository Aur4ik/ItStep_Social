package handler

import (

	"net/http"



	"project/itStep/internal/models"
	"project/itStep/internal/service"
	"project/itStep/internal/utils"
	"project/itStep/internal/dto"
	

	"github.com/gin-gonic/gin"
)
func Register(c *gin.Context) {
    var input dto.RegisterInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
        return
    }

    user := models.User{
        Email:     input.Email,
        Password:  input.Password,
        FirstName: input.FirstName,
        LastName:  input.LastName,
        Group:     input.Group,
    }

    if err := service.Register(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, dto.UserResponse{
        ID:        user.ID,
        Email:     user.Email,
        FirstName: user.FirstName,
        LastName:  user.LastName,
        Group:     user.Group,
        Role:      user.Role,
        Avatar:    user.Avatar,
    })
}

func Login(c *gin.Context) {

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid data",
		})
		return
	}

	user, err := service.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(int(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "token error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func Me(c *gin.Context) {
    userID := c.GetInt("user_id")

    user, err := service.GetUserByID(uint(userID))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    c.JSON(http.StatusOK, dto.ToUserResponse(user))
}