package handlers

import (
	"github.com/dragosv/matrix/commands"
	"github.com/dragosv/matrix/processing"
	"github.com/gin-gonic/gin"
	"net/http"
)

//InvertHandler handler
// @Summary Matrix invert
// @Description returns the matrix as a string in matrix format where the columns and rows are inverted.
// @Tags matrix
// @Accept  mpfd
// @Produce  html
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /invert [post]
func InvertHandler(c *gin.Context) {
	file, err := retrieveFile(c.Request)
	if err != nil {
		writeBadRequestError(c.Writer, err)
		return
	}
	defer file.Close()

	matrix, err := processing.Matrix(file)
	if err != nil {
		writeBadRequestError(c.Writer, err)
		return
	}

	response, err := commands.Invert(matrix)
	if err != nil {
		writeBadRequestError(c.Writer, err)
		return
	}

	c.String(http.StatusOK, response)
}
