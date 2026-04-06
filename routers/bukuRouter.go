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
		bukuRoute.POST("/buku", controllers.CreateBuku)
        bukuRoute.GET("/buku", controllers.GetAllBuku)
        bukuRoute.GET("/buku:id", controllers.GetBukuByID)
        bukuRoute.PUT("/buku:id", controllers.UpdateBuku)
        bukuRoute.DELETE("/buku:id", controllers.DeleteBuku)
	}

	//menambah route Kategori
	kategoriRoute := r.Group("/kategori")
	{
	    kategoriRoute.POST("/kategori", controllers.CreateKategori)
	}

	return r
}