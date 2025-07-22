package routes

import (
	"public_library/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	bookRoutes := router.Group("/books")
	{
		bookRoutes.GET("/", controllers.GetBooks)
		bookRoutes.GET("/:id", controllers.GetBook)
		bookRoutes.POST("/", controllers.CreateBook)
		bookRoutes.PUT("/", controllers.UpdateBook)
		bookRoutes.DELETE("/:id", controllers.DeleteBook)
	}
}
