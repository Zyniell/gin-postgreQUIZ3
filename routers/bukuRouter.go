package routers

import (
	"quizz3-buku/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	// Grouping routes untuk Buku
	bukuRoute := r.Group("/buku")
	{
		bukuRoute.POST("/", controllers.CreateBuku)       // Menjadi: /buku/
        bukuRoute.GET("/", controllers.GetAllBuku)        // Menjadi: /buku/
        bukuRoute.GET("/:id", controllers.GetBukuByID)    // Menjadi: /buku/:id
        bukuRoute.PUT("/:id", controllers.UpdateBuku)     // Menjadi: /buku/:id
        bukuRoute.DELETE("/:id", controllers.DeleteBuku)  // Menjadi: /buku/:id
	}

	//menambah route Kategori
	kategoriRoute := r.Group("/kategori")
	{
	    kategoriRoute.POST("/kategori", controllers.CreateKategori)
	}

	return r
}