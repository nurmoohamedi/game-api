package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurmoohamedi/game-api/db"
	"github.com/nurmoohamedi/game-api/models"
)

// CreateCharacter godoc
// @Summary Создать персонажа для игры
// @Tags characters
// @Accept json
// @Produce json
// @Param id path int true "ID игры"
// @Param character body models.Character true "Данные персонажа"
// @Security BearerAuth
// @Success 201 {object} models.Character
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /games/{id}/characters [post]
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

// GetCharactersByGame godoc
// @Summary Получить всех персонажей для игры
// @Tags characters
// @Produce json
// @Param id path int true "ID игры"
// @Success 200 {array} models.Character
// @Failure 404 {object} map[string]string
// @Router /games/{id}/characters [get]
func GetCharactersByGame(c *gin.Context) {
	var gameID = c.Param("id")

	var characters []models.Character
	db.DB.Where("game_id = ?", gameID).Find(&characters)

	c.IndentedJSON(http.StatusOK, characters)
}

// GetCharacterById godoc
// @Summary Получить персонажа по ID
// @Tags characters
// @Produce json
// @Param id path int true "ID персонажа"
// @Success 200 {object} models.Character
// @Failure 404 {object} map[string]string
// @Router /characters/{id} [get]
func GetCharacterById(c *gin.Context) {
	id := c.Param("id")

	var character models.Character
	if err := db.DB.First(&character, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, character)
}

// UpdateCharacter godoc
// @Summary Обновить персонажа
// @Tags characters
// @Accept json
// @Produce json
// @Param id path int true "ID персонажа"
// @Param character body models.Character true "Данные персонажа"
// @Security BearerAuth
// @Success 200 {object} models.Character
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /characters/{id} [put]
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

// DeleteCharacter godoc
// @Summary Удалить персонажа
// @Tags characters
// @Produce json
// @Param id path int true "ID персонажа"
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /characters/{id} [delete]
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