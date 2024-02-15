package handlers

import (
	"net/http"

	"github.com/JhonatanRian/go-api/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary List Openings
// @Description List all openings
// @Tags openings
// @Accept json
// @Produce json
// @Success 200 {array} schemas.OpeningResponse
// @Router /openings [get]
func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, openings)
}
