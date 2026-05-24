package main


import (
	"github.com/gin-contrib/cors"
	"log"
	"os"

	"project/itStep/internal/config"
	"project/itStep/internal/routes"
	"project/itStep/internal/handler"


	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	
	
err := godotenv.Load()
if err != nil {
    log.Println(".env not found")
}

	config.ConnectDB()

	r := gin.Default()
	r.SetTrustedProxies(nil)

	
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"https://it-step-social-hfxlennyy-auriks-projects.vercel.app",
		},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	routes.SetupRoutes(r)

	
	r.Static("/uploads", "./uploads")
	go handler.HandleMessages()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run("0.0.0.0:" + port)
}