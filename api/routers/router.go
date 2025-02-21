package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// todo make the connection secure with jwt or sessions

func SetupAllRoutes() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	// r.Use(gin.LoggerWithFormatter(middleware.LoggerFormatter))

	// r.Use(middleware.CORSMiddleware())
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
