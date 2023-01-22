package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sumitks866/auth/models"
)

// type UserSignupRequest struct {
// 	Firstname string `json:"firstname" binding:"required"`
// 	Lastname  string `json:"lastname"`
// 	Username  string `json:"username" binding:"required"`
// 	Password  string `json:"password" binding:"required"`
// }

func Signup(ctx *gin.Context) {
	requestBody := models.User{}

	if err := ctx.BindJSON(&requestBody); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := models.AddUser(requestBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, res)
}
