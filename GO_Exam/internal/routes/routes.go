package routes

import (
	"project/itStep/internal/handler"
	"project/itStep/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	auth := r.Group("/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
	}

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleWare())
	{
		api.GET("/me", handler.Me)
	}
}