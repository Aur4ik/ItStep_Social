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
		//user
		api.GET("/me", handler.Me)
		api.POST("/avatar", handler.UploadAvatar)

		//Posts
		api.POST("/posts", handler.CreatePost)
		api.GET("/posts", handler.GetPosts)
		api.DELETE("/posts/:id", handler.DeletePost)
		api.POST("/posts/:id/comments", handler.CreateComment)
		api.GET("/posts/:id/comments", handler.GetComments)
		api.POST("/posts/:id/like", handler.ToggleLike)

		//Comunity
		api.GET("/communities/:id/posts", handler.GetCommunityPosts)
		api.POST("/communities/:id/posts",handler.CreateCommunityPost)
		api.GET("/communities/:id", handler.GetCommunityByID)
		api.GET("/communities/:id/members", handler.GetCommunityMembers)
		api.POST("/communities/:id/join", handler.JoinCommunity)
		api.POST("/communities/:id/leave",handler.LeaveCommunity)
		api.POST("/communities/:id/avatar", handler.CommunityAvatar)
		
	}
	admin := r.Group("/api/admin")

	admin.Use(
		middleware.AuthMiddleWare(),
		middleware.AdminOnly(),
	)
	{
		admin.POST("/users/:id/role", handler.UpdateUserRole)
		admin.DELETE("/communities/:id", handler.DeleteCommunity)
	}

	teacher := api.Group("/communities")

	teacher.Use(middleware.TeacherOrAdmin())
	{
		teacher.POST("/create", handler.CreateCommunity)
		teacher.GET("/all", handler.GetCommunities)
	}
}