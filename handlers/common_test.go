package handlers

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"mime/multipart"
	"net/http"
	"testing"
)

func createRequest(t *testing.T, url string) *http.Request {
	request, err := http.NewRequest(http.MethodPost, url, nil)
	assert.NoError(t, err)

	return request
}

func createFileRequest(t *testing.T, url string, body string) *http.Request {
	buffer := new(bytes.Buffer)
	writer := multipart.NewWriter(buffer)
	formFile, err := writer.CreateFormFile("file", "/path/matrix.csv")
	assert.NoError(t, err)

	bodyBytes := []byte(body)

	_, err = formFile.Write(bodyBytes)
	assert.NoError(t, err)
	assert.NoError(t, writer.Close())

	request, err := http.NewRequest(http.MethodPost, url, buffer)
	assert.NoError(t, err)

	request.Header.Set("Content-type", writer.FormDataContentType())

	return request
}
