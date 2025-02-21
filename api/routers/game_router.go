package routers

import (
	"github.com/Amit152116Kumar/chess_server/api/handlers"
	"github.com/Amit152116Kumar/chess_server/api/websocket"
	"github.com/gin-gonic/gin"
)

func gameRouter(r *gin.Engine) {
	gameGroup := r.Group("/game")
	{
		gameGroup.POST("/create", handlers.CreateGame)
		gameGroup.GET("/:id", handlers.GetGame)

		// gameGroup.GET("/ws/:id", middleware.WSValidationMiddleware(), websocket.WsHandler)
		gameGroup.GET("/ws/:id", websocket.WsHandler)
	}
}
