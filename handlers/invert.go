package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//InvertHandler handler
// @Summary Matrix invert
// @Description returns the matrix as a string in matrix format where the columns and rows are inverted.
// @Tags matrix
// @Accept  mpfd
// @Produce  html
// @Param file formData file true "csv file"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /invert [post]
func (m *MatrixController) InvertHandler(c *gin.Context) {
	file, err := retrieveFile(c.Request)
	if err != nil {
		writeBadRequestError(c.Writer, err)
		return
	}
	defer file.Close()

	matrix, err := m.reader.Read(file)
	if err != nil {
		writeBadRequestError(c.Writer, err)
		return
	}

	response, err := m.operations.Invert(matrix)
	if err != nil {
		writeBadRequestError(c.Writer, err)
		return
	}

	c.String(http.StatusOK, response)
}
