package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurmoohamedi/game-api/db"
	"github.com/nurmoohamedi/game-api/models"
)

// CreateGame godoc
// @Summary Создать игру
// @Tags games
// @Accept json
// @Produce json
// @Param game body models.Game true "Данные игры"
// @Security BearerAuth
// @Success 201 {object} models.Game
// @Failure 400 {object} map[string]string
// @Router /games [post]
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

// GetGames godoc
// @Summary Список игр
// @Tags games
// @Produce json
// @Success 200 {array} models.Game
// @Router /games [get]
func GetGames(c *gin.Context) {
	var games []models.Game
	db.DB.Find(&games)
	c.IndentedJSON(http.StatusOK, games)
}

// GetGameById godoc
// @Summary Получить игру по ID
// @Tags games
// @Produce json
// @Param id path int true "ID игры"
// @Success 200 {object} models.Game
// @Failure 404 {object} map[string]string
// @Router /games/{id} [get]
func GetGameById(c *gin.Context) {
	id := c.Param("id")

	var game models.Game
	if err := db.DB.Preload("Characters").First(&game, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, game)
}

// UpdateGame godoc
// @Summary Обновить игру
// @Tags games
// @Accept json
// @Produce json
// @Param id path int true "ID игры"
// @Param game body models.Game true "Данные игры"
// @Security BearerAuth
// @Success 200 {object} models.Game
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /games/{id} [put]
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

// DeleteGame godoc
// @Summary Удалить игру
// @Tags games
// @Produce json
// @Param id path int true "ID игры"
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /games/{id} [delete]
func DeleteGame(c *gin.Context) {
	id := c.Param("id")

	if err := db.DB.Delete(&models.Game{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete game"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Game deleted successfully"})
}