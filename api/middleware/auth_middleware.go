package middleware

import (
	"github.com/Amit152116Kumar/chess_server/models"
	"github.com/Amit152116Kumar/chess_server/myErrors"
	"github.com/Amit152116Kumar/chess_server/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("session-id")

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": myErrors.SessionMissing.Error()})
			return
		}
		uid, err := uuid.Parse(cookie)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": myErrors.InvalidSession.Error()})
			return
		}
		session, ok := models.Sessions[uid]
		if !ok || !session.IsValid() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": myErrors.SessionExpired.Error()})
			return
		}
		c.Set("session", session)
		c.Next()
	}
}

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := c.MustGet("session").(*models.Session)

		if session.Role == utils.RoleGuest {
			checkPath(c, "/user")
			checkPath(c, "/admin")
		}
		if session.Role == utils.RoleUser {
			checkPath(c, "/admin")
		}
		c.Next()
	}
}

func checkPath(c *gin.Context, subPath string) {
	ok := strings.Contains(c.Request.URL.Path, subPath)
	if ok {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": myErrors.Forbidden.Error()})
	}
	return
}

func WSValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		_, err := uuid.Parse(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

	}
}
