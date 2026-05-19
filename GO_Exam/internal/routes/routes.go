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
		// user
		api.GET("/me", handler.Me)
		api.POST("/avatar", handler.UploadAvatar)

		// posts
		api.POST("/posts", handler.CreatePost)
		api.GET("/posts", handler.GetPosts)
		api.DELETE("/posts/:id", handler.DeletePost)
		api.GET("/posts/:id", handler.GetPostByID) // было /post/:id — исправлено на /posts/:id

		api.POST("/posts/:id/comments", handler.CreateComment)
		api.GET("/posts/:id/comments", handler.GetComments)

		api.POST("/posts/:id/like", handler.ToggleLike)

		// communities
		api.GET("/communities/:id/posts", handler.GetCommunityPosts)
		api.POST("/communities/:id/posts", handler.CreateCommunityPost)
		api.GET("/communities/:id", handler.GetCommunityByID)
		api.GET("/communities/:id/members", handler.GetCommunityMembers)
		api.POST("/communities/:id/join", handler.JoinCommunity)
		api.POST("/communities/:id/leave", handler.LeaveCommunity)
		api.POST("/communities/:id/avatar", handler.CommunityAvatar)
		api.GET("/communities/:id/schedule", handler.GetCommunitySchedule)

		// chats
		api.POST("/chats", handler.CreateChat)
		api.GET("/chats", handler.GetChats)
		api.PUT("/chats/:id", handler.UpdateChat)
		api.DELETE("/chats/:id", handler.DeleteChat)

		api.POST("/chats/:id/members", handler.AddChatMember)
		api.GET("/chats/:id/members", handler.GetChatMembers)

		api.POST("/chats/:id/messages", handler.CreateMessage)
		api.GET("/chats/:id/messages", handler.GetChatMessages)

		api.POST("/dm", handler.CreateDMChat)

		// WebSocket — теперь за AuthMiddleWare
		api.GET("/ws", handler.HandleConnections)
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

	teacher := api.Group("/teacher")
	teacher.Use(middleware.TeacherOrAdmin())
	{
		teacher.POST("/communities/create", handler.CreateCommunity)
		teacher.GET("/communities/all", handler.GetCommunities)
		teacher.POST("/communities/:id/schedule", handler.CreateSchedule)
		teacher.PUT("/schedule/:id/update", handler.UpdateSchedule)
		teacher.DELETE("/schedule/:id/delete", handler.DeleteSchedule)
	}
}
