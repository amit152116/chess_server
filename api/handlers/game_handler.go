package handlers

import (
	"fmt"
	"net/http"

	"github.com/Amit152116Kumar/chess_server/models"
	"github.com/Amit152116Kumar/chess_server/services"
	"github.com/gin-gonic/gin"
)

func CreateGame(c *gin.Context) {
	fmt.Println(c.ContentType())
	var gameTimeControl *models.NewGameReqParam
	if err := c.ShouldBind(&gameTimeControl); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid := services.CreateGame(gameTimeControl)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Random game created", "game_id": uid})
}

func GetGame(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Game found"})
}
