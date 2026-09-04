package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurmoohamedi/game-api/db"
	"github.com/nurmoohamedi/game-api/models"
)

func CreateGame(c *gin.Context) {
	var game models.Game
	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&game).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create game"})
		return
	}

	c.IndentedJSON(http.StatusCreated, game)
}

func GetGames(c *gin.Context) {
	var games []models.Game
	db.DB.Find(&games)
	c.IndentedJSON(http.StatusOK, games)
}

func GetGameById(c *gin.Context) {
	id := c.Param("id")

	var game models.Game
	if err := db.DB.Preload("Characters").First(&game, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, game)
}

func UpdateGame(c *gin.Context) {
	id := c.Param("id")

	var game models.Game
	if err := db.DB.First(&game, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&game)
	c.IndentedJSON(http.StatusOK, game)
}

func DeleteGame(c *gin.Context) {
	id := c.Param("id")

	if err := db.DB.Delete(&models.Game{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete game"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Game deleted successfully"})
}