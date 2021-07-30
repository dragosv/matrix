package handlers

import (
	"fmt"
	"github.com/dragosv/matrix/commands"
	"github.com/dragosv/matrix/processing"
	"net/http"
)

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	file, err := retrieveFile(r)
	if err != nil {
		writeError(w, err)
		return
	}
	defer file.Close()

	records, err := processing.Records(file)
	if err != nil {
		writeError(w, err)
		return
	}

	response, err := commands.Multiply(records)
	if err != nil {
		writeError(w, err)
		return
	}

	fmt.Fprint(w, response)
}
