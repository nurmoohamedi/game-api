package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nurmoohamedi/game-api/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	games := router.Group("/games")
	{
		games.POST("", handlers.CreateGame)
		games.GET("", handlers.GetGames)
		games.GET("/:id", handlers.GetGameById)
		games.PUT("/:id", handlers.UpdateGame)
		games.DELETE("/:id", handlers.DeleteGame)
		// Character routes nested under games
		games.POST("/:id/characters", handlers.CreateCharacter)
		games.GET("/:id/characters", handlers.GetCharactersByGame)
	}

	characters := router.Group("/characters")
	{
		characters.GET("/:id", handlers.GetCharacterById)
		characters.PUT("/:id", handlers.UpdateCharacter)
		characters.DELETE("/:id", handlers.DeleteCharacter)
	}

	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
}