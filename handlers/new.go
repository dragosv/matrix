package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//NewHandler handler
// @Summary Matrix echo
// @Description Returns the matrix as a string in matrix format
// @Tags matrix
// @Accept  mpfd
// @Produce  html
// @Param file formData file true "csv file"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /echo [post]
func (m *MatrixController) NewHandler(c *gin.Context) {
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

	response, err := m.operations.New(matrix)
	if err != nil {
		writeBadRequestError(c.Writer, err)
		return
	}

	c.String(http.StatusOK, response)
}
