package main

import (
	"log"

	"project/itStep/internal/config"
	"project/itStep/internal/routes"

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

	routes.SetupRoutes(r)

	r.Run(":8080")
}