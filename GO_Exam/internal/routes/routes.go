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
		api.POST("/avatar", handler.UploadAvatar)

		api.POST("/posts", handler.CreatePost)
		api.GET("posts", handler.GetPosts)
		api.DELETE("posts/:id", handler.DeletePost)
		api.POST("/posts/:id/comments", handler.CreateComment)
		api.GET("/posts/:id/comments", handler.GetComments)
	}
}