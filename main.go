package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/nurmoohamedi/game-api/db"
	"github.com/nurmoohamedi/game-api/models"
	"github.com/nurmoohamedi/game-api/routes"
)

// @title Game API
// @version 1.0
// @description REST API для управления играми и персонажами
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db.Connect()

	if err := db.DB.AutoMigrate(&models.Game{}, &models.Character{}, &models.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router)
	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.Run(":8080")
}