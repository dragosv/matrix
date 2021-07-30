package handlers

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"mime/multipart"
	"net/http"
	"testing"
)

func createRequest(t *testing.T, body string) *http.Request {
	buffer := new(bytes.Buffer)
	writer := multipart.NewWriter(buffer)
	formFile, err := writer.CreateFormFile("file", "/path/matrix.csv")
	assert.NoError(t, err)

	bodyBytes := []byte(body)

	_, err = formFile.Write(bodyBytes)
	assert.NoError(t, err)
	assert.NoError(t, writer.Close())

	request, err := http.NewRequest(http.MethodPost, "/echo", buffer)
	assert.NoError(t, err)

	request.Header.Set("Content-type", writer.FormDataContentType())

	return request
}
