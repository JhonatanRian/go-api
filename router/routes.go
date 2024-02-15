package router

import (
	"github.com/JhonatanRian/go-api/handlers"
	"github.com/gin-gonic/gin"
)

func initializeOpeningRoutes(router *gin.Engine) {
	handlers.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/openings", handlers.CreateOpening)
		// v1.POST("/openings", openinghandlers.CreateOpening)
		// v1.GET("/openings", openinghandlers.CreateOpening)
		// v1.PUT("/openings", openinghandlers.CreateOpening)
		// v1.DELETE("/openings", openinghandlers.CreateOpening)
	}
}
