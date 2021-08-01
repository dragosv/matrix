package handlers

import (
	"github.com/dragosv/matrix/commands"
	"github.com/dragosv/matrix/processing"
	"github.com/gin-gonic/gin"
	"net/http"
)

//FlattenHandler handler
// @Summary Matrix flatten
// @Description Returns the matrix as a 1 line string, with values separated by commas.
// @Tags matrix
// @Accept  mpfd
// @Produce  html
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /flatten [post]
func FlattenHandler(c *gin.Context) {
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

	response, err := commands.Flatten(matrix)
	if err != nil {
		writeBadRequestError(c.Writer, err)
		return
	}

	c.String(http.StatusOK, response)
}
