package main

import (
	"os"
	"quizz3-buku/config"
	"quizz3-buku/routers"


)

func main() {
	// 1. Koneksi ke Database
	config.ConnectDB()

	// 2. Jalankan Server
	r := routers.StartServer()
	
	r.Run(":" + os.Getenv("PORT"))
}