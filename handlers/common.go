package handlers

import (
	"fmt"
	"mime/multipart"
	"net/http"
)

func retrieveFile(r *http.Request) (multipart.File, error) {
	file, _, err := r.FormFile("file")
	return file, err
}

func writeBadRequestError(w http.ResponseWriter, err error) (int, error) {
	return writeError(w, err, http.StatusBadRequest)
}

func writeError(w http.ResponseWriter, err error, statusCode int) (int, error) {
	w.WriteHeader(statusCode)
	return w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
}
