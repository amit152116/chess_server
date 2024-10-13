package handlers

import (
	"fmt"
	"github.com/Amit152116Kumar/chess_server/models"
	"github.com/Amit152116Kumar/chess_server/services"
	"github.com/Amit152116Kumar/chess_server/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func Login(c *gin.Context) {
	var credentials models.LoginUserPayload

	if err := c.ShouldBind(&credentials); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, err := services.AuthenticateUser(&credentials)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error(), "message": "Email or password is incorrect"})
		return
	}
	writeHeader(c, uid.String())
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User authenticated successfully"})
}

func Register(c *gin.Context) {

	var user models.RegisterUserPayload

	if err := c.ShouldBind(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, err := services.RegisterUser(&user)
	if err != nil {
		c.IndentedJSON(http.StatusConflict, gin.H{"error": err.Error(), "message": "User already exists", "user": user})
		return
	}
	writeHeader(c, uid.String())
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func writeHeader(c *gin.Context, sessionId string) {
	c.SetCookie("session-id", sessionId, int(utils.SessionTimeout), "/", "", false, true)
	c.Header("expiry", fmt.Sprintf("%d", utils.SessionTimeout))
}

func Guest(c *gin.Context) {
	uid := services.Guest()
	writeHeader(c, uid.String())
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Guest user created"})
}

func RefreshToken(c *gin.Context) {
	cookie, _ := c.Cookie("session-id")
	newUid := services.RefreshToken(cookie)
	writeHeader(c, newUid.String())
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Token refreshed successfully"})
}

func Logout(c *gin.Context) {
	sessionId := c.Request.Header.Get("session-id")
	uid := uuid.MustParse(sessionId)
	delete(models.Sessions, uid)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
}
