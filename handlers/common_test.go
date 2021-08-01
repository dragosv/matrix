package handlers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func matrixRequest(t *testing.T, url string) (string, int) {
	gin.SetMode(gin.TestMode)

	request := createRequest(t, url)
	response := httptest.NewRecorder()

	_, engine := gin.CreateTestContext(response)

	SetupMatrixHandlers(engine)

	engine.ServeHTTP(response, request)

	responseBody := response.Body.String()
	code := response.Code
	return responseBody, code
}

func matrixFileRequest(t *testing.T, url string, body string) (string, int) {
	gin.SetMode(gin.TestMode)

	request := createFileRequest(t, url, body)
	response := httptest.NewRecorder()

	_, engine := gin.CreateTestContext(response)

	SetupMatrixHandlers(engine)

	engine.ServeHTTP(response, request)

	responseBody := response.Body.String()
	code := response.Code

	return responseBody, code
}

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
