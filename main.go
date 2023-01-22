package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sumitks866/auth/models"
	router "github.com/sumitks866/auth/routers"
)

func main() {
	fmt.Print("Server starting...")
	r := router.InitializeRouter()
	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	models.Setup()

	log.Fatal(server.ListenAndServe())
}
