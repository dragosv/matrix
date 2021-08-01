package handlers

import (
	"github.com/dragosv/matrix/commands"
	"github.com/dragosv/matrix/processing"
	"github.com/gin-gonic/gin"
	"net/http"
)

//EchoHandler handler
// @Summary Matrix echo
// @Description Returns the matrix as a string in matrix format
// @Tags matrix
// @Accept  mpfd
// @Produce  html
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /echo [post]
func EchoHandler(c *gin.Context) {
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

	response, err := commands.Echo(matrix)
	if err != nil {
		writeBadRequestError(c.Writer, err)
		return
	}

	c.String(http.StatusOK, response)
}
