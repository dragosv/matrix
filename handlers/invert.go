package handlers

import (
	"fmt"
	"github.com/dragosv/matrix/commands"
	"github.com/dragosv/matrix/processing"
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
func InvertHandler(w http.ResponseWriter, r *http.Request) {
	file, err := retrieveFile(r)
	if err != nil {
		writeBadRequestError(w, err)
		return
	}
	defer file.Close()

	matrix, err := processing.Matrix(file)
	if err != nil {
		writeBadRequestError(w, err)
		return
	}

	response, err := commands.Invert(matrix)
	if err != nil {
		writeBadRequestError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, response)
}
