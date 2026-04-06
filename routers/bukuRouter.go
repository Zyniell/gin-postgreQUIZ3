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
		r.POST("/buku", controllers.CreateBuku)
		r.GET("/buku", controllers.GetAllBuku)
		r.GET("/buku/:id", controllers.GetBukuByID)
		r.PUT("/buku/:id", controllers.UpdateBuku)
		r.DELETE("/buku/:id", controllers.DeleteBuku)
	}

	//menambah route Kategori
	kategoriRoute := r.Group("/kategori")
	{
	    kategoriRoute.POST("/", controllers.CreateKategori)
	}

	return r
}