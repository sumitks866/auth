package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sumitks866/auth/models"
)

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(ctx *gin.Context) {
	requestBody := UserLoginRequest{}

	if err := ctx.BindJSON(&requestBody); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := models.GetUser(requestBody.Username, requestBody.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	ctx.JSON(http.StatusOK, res)

}
