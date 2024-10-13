package routers

import (
	"github.com/Amit152116Kumar/chess_server/api/handlers"
	"github.com/Amit152116Kumar/chess_server/api/middleware"
	"github.com/gin-gonic/gin"
)

func userRouter(r *gin.Engine) {
	userGroup := r.Group("/user", middleware.AuthenticationMiddleware(), middleware.AuthorizationMiddleware())
	{
		userGroup.GET("", handlers.GetUser)
		userGroup.PUT("", handlers.UpdateUser)
		userGroup.DELETE("", handlers.DeleteUser)
		userGroup.GET("/games", handlers.GetGamesForUser)

	}

}
