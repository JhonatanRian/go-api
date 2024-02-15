package handlers

import (
	"net/http"

	"github.com/JhonatanRian/go-api/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary Get Opening Details
// @Description Retrieve details of an opening by ID
// @Tags openings
// @Accept json
// @Produce json
// @Param id path string true "Opening ID"
// @Success 200 {object} schemas.OpeningResponse
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /openings/{id} [get]
func DetailOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, opening)
}
