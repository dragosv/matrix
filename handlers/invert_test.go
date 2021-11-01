package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestInvert_WithFile_ShouldOutputExpected(t *testing.T) {
	body := "1,2,3\n4,5,6\n7,8,9"
	url := "/invert"

	responseBody, code := matrixFileRequest(t, url, body)

	assert.Equal(t, "1,4,7\n2,5,8\n3,6,9", responseBody)
	assert.Equal(t, http.StatusOK, code)
}

func TestInvert_WithoutFile_ShouldOutputError(t *testing.T) {
	url := "/invert"

	responseBody, code := matrixRequest(t, url)

	assert.Equal(t, "error request Content-Type isn't multipart/form-data", responseBody)
	assert.Equal(t, http.StatusBadRequest, code)
}

func TestInvert_WithNonNumbers_ShouldOutputError(t *testing.T) {
	body := "a"
	url := "/invert"

	responseBody, code := matrixFileRequest(t, url, body)

	assert.Equal(t, "error all values should be numbers", responseBody)
	assert.Equal(t, http.StatusBadRequest, code)
}

func TestInvert_WithNonCsv_ShouldOutputError(t *testing.T) {
	body := "1,2,3\n4,5,6\n7,8"
	url := "/invert"

	responseBody, code := matrixFileRequest(t, url, body)

	assert.Equal(t, "error could not read the file", responseBody)
	assert.Equal(t, http.StatusBadRequest, code)
}

func TestInvert_WithNonMatrix_ShouldOutputError(t *testing.T) {
	body := "1,2,3\n4,5,6\n7,8,9\n10,11,12"
	url := "/invert"

	responseBody, code := matrixFileRequest(t, url, body)

	assert.Equal(t, "error the number of columns and rows should be the same", responseBody)
	assert.Equal(t, http.StatusBadRequest, code)
}
