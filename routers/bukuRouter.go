package routes

import (
	"quizz3-buku/controllers"

	"github.com/gin-gonic/gin"
)

func BukuRoutes(router *gin.Engine) {
	route := router.Group("/buku")
	{
		route.GET("/", controllers.GetBuku)      // Get All
		route.GET("/:id", controllers.GetBuku)   // Get by ID
		route.POST("/", controllers.CreateBuku)
		route.PUT("/:id", controllers.UpdateBuku)
		route.DELETE("/:id", controllers.DeleteBuku)
	}
}