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
		bukuRoute.POST("", controllers.CreateBuku)       
		bukuRoute.GET("", controllers.GetAllBuku)        
		bukuRoute.GET("/:id", controllers.GetBukuByID)    
		bukuRoute.PUT("/:id", controllers.UpdateBuku)     
		bukuRoute.DELETE("/:id", controllers.DeleteBuku)
	}

	//menambah route Kategori
	kategoriRoute := r.Group("/kategori")
	{
	    kategoriRoute.POST("", controllers.CreateKategori)
	}

	return r
}