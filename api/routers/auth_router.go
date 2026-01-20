package routers

import (
	"github.com/amit152116/chess_server/api/handlers"
	"github.com/amit152116/chess_server/api/middleware"
	"github.com/gin-gonic/gin"
)

func authRouter(r *gin.Engine) {
	authGroup := r.Group("/")
	{
		authGroup.POST("/login", handlers.Login)
		authGroup.POST("/register", handlers.Register)
		authGroup.GET("/guest", handlers.Guest)
		authGroup.GET("/logout", middleware.AuthorizationMiddleware(), handlers.Logout)
		authGroup.GET("/refresh", middleware.AuthenticationMiddleware(), handlers.RefreshToken)

	}
}
