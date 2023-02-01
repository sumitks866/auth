package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sumitks866/auth/middleware/cors"
	"github.com/sumitks866/auth/pkg/auth"
	"github.com/sumitks866/auth/routers/api"
)

func InitializeRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.CORSMiddleware())

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	router.GET("/verify", func(ctx *gin.Context) {
		fmt.Println("verifying...")
		token := ctx.Query("token")
		res, err := auth.ParseToken(token)

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, res)
	})
	router.POST("/login", api.Login)
	router.POST("/signup", api.Signup)
	return router
}
