package main

import (
	"os"
	"quizz3-buku/config"
	"quizz3-buku/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi Database
	config.ConnectDB()

	router := gin.Default()

	// Panggil Router
	routes.BukuRoutes(router)

	// Port untuk Railway
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}