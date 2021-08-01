package handlers

import (
	"github.com/dragosv/matrix/commands"
	"github.com/dragosv/matrix/processing"
	"github.com/gin-gonic/gin"
	"net/http"
)

//SumHandler handler
// @Summary Matrix sum
// @Description Returns the sum of the integers in the matrix.
// @Tags matrix
// @Accept  mpfd
// @Produce  html
// @Param file formData file true "csv file"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /sum [post]
func SumHandler(c *gin.Context) {
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

	response, err := commands.Sum(matrix)
	if err != nil {
		writeBadRequestError(c.Writer, err)
		return
	}

	c.String(http.StatusOK, response)
}
