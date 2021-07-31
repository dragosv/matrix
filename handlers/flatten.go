package handlers

import (
	"fmt"
	"github.com/dragosv/matrix/commands"
	"github.com/dragosv/matrix/processing"
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
func FlattenHandler(w http.ResponseWriter, r *http.Request) {
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

	response, err := commands.Flatten(matrix)
	if err != nil {
		writeBadRequestError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, response)
}
