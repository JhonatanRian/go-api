package handlers

import (
	"net/http"

	"github.com/JhonatanRian/go-api/schemas"
	"github.com/gin-gonic/gin"
)

func DetailOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, opening)
}