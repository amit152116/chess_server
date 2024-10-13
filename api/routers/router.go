package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// todo make the connection secure with jwt or sessions

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//r.Use(gin.LoggerWithFormatter(middleware.LoggerFormatter))

	homeRouter(r)
	userRouter(r)
	gameRouter(r)
	authRouter(r)
	adminRoute(r)

	return r
}

func homeRouter(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Welcome to chess server"})
	})
}
