package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFlatten_WithValidMatrix_ShouldOutputExpected(t *testing.T) {
	body := "1,2,3\n4,5,6\n7,8,9"

	request := createFileRequest(t, "/echo", body)
	response := httptest.NewRecorder()

	FlattenHandler(response, request)

	responseBody := response.Body.String()

	assert.Equal(t, "1,2,3,4,5,6,7,8,9", responseBody)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestFlatten_WithNonNumbers_ShouldOutputError(t *testing.T) {
	body := "a"

	request := createFileRequest(t, "/echo", body)
	response := httptest.NewRecorder()

	FlattenHandler(response, request)

	responseBody := response.Body.String()

	assert.Equal(t, "error all values should be numbers", responseBody)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestFlatten_WithNonCsv_ShouldOutputError(t *testing.T) {
	body := "1,2,3\n4,5,6\n7,8"

	request := createFileRequest(t, "/echo", body)
	response := httptest.NewRecorder()

	FlattenHandler(response, request)

	responseBody := response.Body.String()

	assert.Equal(t, "error could not read the file", responseBody)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func TestFlatten_WithNonMatrix_ShouldOutputError(t *testing.T) {
	body := "1,2,3\n4,5,6\n7,8,9\n10,11,12"

	request := createFileRequest(t, "/echo", body)
	response := httptest.NewRecorder()

	FlattenHandler(response, request)

	responseBody := response.Body.String()

	assert.Equal(t, "error the number of columns and rows should be the same", responseBody)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}
