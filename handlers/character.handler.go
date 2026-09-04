package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurmoohamedi/game-api/db"
	"github.com/nurmoohamedi/game-api/models"
)

func CreateCharacter(c *gin.Context) {
	gameID := c.Param("id")

	var game models.Game
	if err := db.DB.First(&game, gameID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}

	var character models.Character
	if err := c.ShouldBindJSON(&character); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	character.GameID = game.ID
	if err := db.DB.Create(&character).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create character"})
		return
	}

	c.IndentedJSON(http.StatusCreated, character)
}

func GetCharactersByGame(c *gin.Context) {
	var gameID = c.Param("id")

	var characters []models.Character
	db.DB.Where("game_id = ?", gameID).Find(&characters)

	c.IndentedJSON(http.StatusOK, characters)
}

func GetCharacterById(c *gin.Context) {
	id := c.Param("id")

	var character models.Character
	if err := db.DB.First(&character, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, character)
}

func UpdateCharacter(c *gin.Context) {
	id := c.Param("id")

	var character models.Character
	if err := db.DB.First(&character, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
		return
	}

	if err := c.ShouldBindJSON(&character); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&character)
	c.IndentedJSON(http.StatusOK, character)
}

func DeleteCharacter(c *gin.Context) {
	id := c.Param("id")

	var character models.Character
	if err := db.DB.First(&character, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
		return
	}

	db.DB.Delete(&character)
	c.Status(http.StatusNoContent)
}