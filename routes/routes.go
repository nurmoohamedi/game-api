package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/nurmoohamedi/game-api/handlers"
	"github.com/nurmoohamedi/game-api/middleware"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	// Публичные — чтение доступно всем
	router.GET("/games", handlers.GetGames)
	router.GET("/games/:id", handlers.GetGameById)
	router.GET("/games/:id/characters", handlers.GetCharactersByGame)
	router.GET("/characters/:id", handlers.GetCharacterById)

	// Приватные — требуют аутентификации
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/games", handlers.CreateGame)
		protected.PUT("/games/:id", handlers.UpdateGame)
		protected.DELETE("/games/:id", handlers.DeleteGame)

		protected.POST("/games/:id/characters", handlers.CreateCharacter)
		protected.PUT("/characters/:id", handlers.UpdateCharacter)
		protected.DELETE("/characters/:id", handlers.DeleteCharacter)
	}
}