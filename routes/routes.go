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
	}
}