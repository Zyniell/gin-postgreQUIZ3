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
		bukuRoute.POST("/", controllers.CreateBuku)       // Tambah buku baru
		bukuRoute.GET("/", controllers.GetAllBuku)       // Ambil semua daftar buku
		bukuRoute.GET("/:id", controllers.GetBukuByID)   // Ambil satu buku berdasarkan ID
		bukuRoute.PUT("/:id", controllers.UpdateBuku)    // Update data buku
		bukuRoute.DELETE("/:id", controllers.DeleteBuku) // Hapus buku
	}

	//menambah route Kategori
	kategoriRoute := r.Group("/kategori")
	{
	    kategoriRoute.POST("/", controllers.CreateKategori)
	}

	return r
}