package routers

import (
	"github.com/Amit152116Kumar/chess_server/api/handlers"
	"github.com/Amit152116Kumar/chess_server/api/middleware"
	"github.com/gin-gonic/gin"
)

func adminRoute(r *gin.Engine) {
	adminGroup := r.Group("/admin", middleware.AuthenticationMiddleware(), middleware.AuthorizationMiddleware())
	{
		adminGroup.GET("", handlers.GetUser)
		adminGroup.PUT("", handlers.UpdateUser)
		adminGroup.DELETE("", handlers.DeleteUser)
		adminGroup.GET("/games", handlers.GetGamesForUser)
		adminGroup.GET("/logout", handlers.Logout)
		adminGroup.GET("/refresh", handlers.RefreshToken)
	}

}
