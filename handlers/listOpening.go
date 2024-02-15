package handlers

import (
	"net/http"

	"github.com/JhonatanRian/go-api/schemas"
	"github.com/gin-gonic/gin"
)


func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, openings)
}
