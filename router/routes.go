package router

import (
	"github.com/JhonatanRian/go-api/handlers"
	"github.com/gin-gonic/gin"
	docs "github.com/JhonatanRian/go-api/docs"
	swaggerfiles "github.com/swaggo/files"
   	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeOpeningRoutes(router *gin.Engine) {
	handlers.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/openings", handlers.ListOpeningHandler)
		v1.POST("/openings", handlers.CreateOpeningHandler)
		v1.GET("/openings/:id", handlers.DetailOpeningHandler)
		v1.PUT("/openings/:id", handlers.UpdateOpeningHandler)
		v1.DELETE("/openings/:id", handlers.DeleteOpeningHandler)
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
