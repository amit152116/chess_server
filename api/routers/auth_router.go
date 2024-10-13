package routers

import (
	"github.com/Amit152116Kumar/chess_server/api/handlers"
	"github.com/Amit152116Kumar/chess_server/api/middleware"
	"github.com/gin-gonic/gin"
)

func authRouter(r *gin.Engine) {
	authGroup := r.Group("/", middleware.RateLimiter())
	{
		authGroup.POST("/login", handlers.Login)
		authGroup.POST("/register", handlers.Register)
		authGroup.GET("/guest", handlers.Guest)
		authGroup.GET("/logout", handlers.Logout)
		authGroup.GET("/refresh", middleware.AuthenticationMiddleware(), handlers.RefreshToken)

	}
}
