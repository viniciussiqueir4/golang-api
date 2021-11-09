package routes

import (
	"example/webapi/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.POST("/", controllers.CreateBook)
			books.GET("/:id", controllers.ShowBook)
		}
	}

	return router
}
