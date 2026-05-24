package main


import (
	"github.com/gin-contrib/cors"
	"log"

	"project/itStep/internal/config"
	"project/itStep/internal/routes"
	"project/itStep/internal/handler"


	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env not loaded")
	}

	config.ConnectDB()

	r := gin.Default()

	
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	routes.SetupRoutes(r)

	
	r.Static("/uploads", "./uploads")
	go handler.HandleMessages()
	r.Run(":8080")
}