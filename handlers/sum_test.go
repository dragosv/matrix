package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestSum_WithValidMatrix_ShouldOutputExpected(t *testing.T) {
	body := "1,2,3\n4,5,6\n7,8,9"

	request := createFileRequest(t, "/echo", body)
	response := httptest.NewRecorder()

	SumHandler(response, request)

	responseBody := response.Body.String()

	assert.Equal(t, "45", responseBody)
}

func TestSum_WithNonNumbers_ShouldOutputError(t *testing.T) {
	body := "a"

	request := createFileRequest(t, "/echo", body)
	response := httptest.NewRecorder()

	SumHandler(response, request)

	responseBody := response.Body.String()

	assert.Equal(t, "error all values should be numbers", responseBody)
}

func TestSum_WithNonCsv_ShouldOutputError(t *testing.T) {
	body := "1,2,3\n4,5,6\n7,8"

	request := createFileRequest(t, "/echo", body)
	response := httptest.NewRecorder()

	SumHandler(response, request)

	responseBody := response.Body.String()

	assert.Equal(t, "error could not read the file", responseBody)
}

func TestSum_WithNonMatrix_ShouldOutputError(t *testing.T) {
	body := "1,2,3\n4,5,6\n7,8,9\n10,11,12"

	request := createFileRequest(t, "/echo", body)
	response := httptest.NewRecorder()

	SumHandler(response, request)

	responseBody := response.Body.String()

	assert.Equal(t, "error the number of columns and rows should be the same", responseBody)
}
