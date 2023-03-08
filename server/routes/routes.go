package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/lucasdk3/maui-oqcomer-api/controllers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// "github.com/lucasdk3/maui-oqcomer-appserver/middlewares"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("user")
		{
			user.POST("/", controllers.CreateUser)
			user.GET("/", controllers.GetUser)
		}
		// books := main.Group("books", middlewares.Auth())
		// {
		// 	books.GET("/", controllers.ShowAllBooks)
		// 	books.GET("/:id", controllers.ShowBook)
		// 	books.POST("/", controllers.CreateBook)
		// 	books.PUT("/", controllers.EditBook)
		// 	books.DELETE("/:id", controllers.DeleteBook)
		// }

		main.POST("login", controllers.Login)

		main.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return router
}
