package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/sumitks866/auth/models"
	"github.com/sumitks866/auth/pkg/auth"
	router "github.com/sumitks866/auth/routers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
	fmt.Print("Server starting...")
	r := router.InitializeRouter()
	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	auth.Setup()
	models.Setup()

	log.Fatal(server.ListenAndServe())
}
