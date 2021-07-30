package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestInvert_WithFile_ShouldOutputExpected(t *testing.T) {
	body := "1,2,3\n4,5,6\n7,8,9"

	request := createRequest(t, body)
	response := httptest.NewRecorder()

	InvertHandler(response, request)

	responseBody := response.Body.String()

	assert.Equal(t, "1,4,7\n2,5,8\n3,6,9", responseBody)
}

func TestInvert_WithNonNumbers_ShouldOutputError(t *testing.T) {
	body := "a"

	request := createRequest(t, body)
	response := httptest.NewRecorder()

	InvertHandler(response, request)

	responseBody := response.Body.String()

	assert.Equal(t, "error all values should be numbers", responseBody)
}

func TestInvert_WithNonCsv_ShouldOutputError(t *testing.T) {
	body := "1,2,3\n4,5,6\n7,8"

	request := createRequest(t, body)
	response := httptest.NewRecorder()

	InvertHandler(response, request)

	responseBody := response.Body.String()

	assert.Equal(t, "error could not read the file", responseBody)
}

func TestInvert_WithNonMatrix_ShouldOutputError(t *testing.T) {
	body := "1,2,3\n4,5,6\n7,8,9\n10,11,12"

	request := createRequest(t, body)
	response := httptest.NewRecorder()

	InvertHandler(response, request)

	responseBody := response.Body.String()

	assert.Equal(t, "error the number of columns and rows should be the same", responseBody)
}
