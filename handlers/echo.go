package handlers

import (
	"fmt"
	"github.com/dragosv/matrix/commands"
	"github.com/dragosv/matrix/processing"
	"net/http"
)

//EchoHandler processes the '/echo' endpoint
func EchoHandler(w http.ResponseWriter, r *http.Request) {
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

	response, err := commands.Echo(matrix)
	if err != nil {
		writeBadRequestError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, response)
}
