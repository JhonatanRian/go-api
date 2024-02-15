package router

import "github.com/gin-gonic/gin"

func Initialize() {
	r := gin.Default()
	initializeOpeningRoutes(r)	
	r.Run()
}
