package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/lucasdk3/go-oqcomer-api/controllers"
	"github.com/lucasdk3/go-oqcomer-api/controllers/categories"
	"github.com/lucasdk3/go-oqcomer-api/controllers/foods"
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
		categories_routes := main.Group("categories")
		{
			categories_routes.GET("/", categories.GetCategories)
			categories_routes.GET("/:id", categories.GetCategory)
			categories_routes.POST("/", categories.AddCategory)
			categories_routes.PUT("/", categories.EditCategory)
			categories_routes.DELETE("/:id", categories.DeleteCategory)
		}

		foods_routes := main.Group("foods")
		{
			foods_routes.GET("/", foods.GetFoods)
			foods_routes.GET("/:id", foods.GetFood)
			foods_routes.GET("/sort/:id", foods.SortFood)
			foods_routes.POST("/", foods.AddFood)
			foods_routes.PUT("/", foods.EditFood)
			foods_routes.DELETE("/:id", foods.DeleteFood)
		}

		main.POST("login", controllers.Login)

		main.POST("refresh", controllers.Refresh)

		main.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return router
}
