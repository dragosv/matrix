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

func writeError(w http.ResponseWriter, err error) (int, error) {
	return w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
}
