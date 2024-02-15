package handlers

import (
	"net/http"

	"github.com/JhonatanRian/go-api/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary Create Opening
// @Description Create a new opening
// @Tags openings
// @Accept json
// @Produce json
// @Param opening body CreateOpeningRequest true "Opening"
// @Success 201 {object} schemas.OpeningResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /openings [post]
func CreateOpeningHandler(c *gin.Context) {
	request := CreateOpeningRequest{}
	c.BindJSON(&request)
	request.Validate()

	if err := request.Validate(); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, opening)
}
