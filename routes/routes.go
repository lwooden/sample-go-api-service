package routes

import (
	"sample-go-api-service/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", handlers.HealthCheck)

	v1 := router.Group("/v1")
	{
		v1.GET("api/private", handlers.GenerateMessage)
		v1.GET("api/public", handlers.FetchCatFacts)
		// 	v1.GET("api/category", Handlers.GetAllCategories)
		// 	v1.POST("api/category", Handlers.CreateCategory)

		// 	// v1.GET("todo/:id", Controllers.GetATodo)
		// 	// v1.PUT("todo/:id", Controllers.UpdateATodo)
		// 	// v1.DELETE("todo/:id", Controllers.DeleteATodo)
	}
	return router
}
