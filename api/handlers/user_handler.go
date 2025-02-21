package handlers

import (
	"github.com/Amit152116Kumar/chess_server/models"
	"github.com/Amit152116Kumar/chess_server/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateUser(&user); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "DBUser updated", "user": user})

}

func GetUser(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, gin.H{"message": "DBUser found"})
}

func GetGamesForUser(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Games found for user"})
}

func DeleteUser(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, gin.H{"message": "DBUser deleted"})
}
