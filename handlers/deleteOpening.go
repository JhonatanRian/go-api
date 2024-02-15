package handlers

import (
	"net/http"

	"github.com/JhonatanRian/go-api/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary Delete Opening
// @Description Delete an opening by ID
// @Tags openings
// @Accept json
// @Produce json
// @Param id path string true "Opening ID"
// @Success 204 "No Content"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /openings/{id} [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, err.Error())
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.Status(http.StatusNoContent)
}
