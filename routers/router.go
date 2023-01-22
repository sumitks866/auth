package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sumitks866/auth/routers/api"
)

func InitializeRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	router.POST("/login", api.Login)
	router.POST("/signup", api.Signup)
	return router
}
