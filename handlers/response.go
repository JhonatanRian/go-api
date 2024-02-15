package handlers

import "github.com/gin-gonic/gin"

func sendError(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, gin.H{"error": message})
}
