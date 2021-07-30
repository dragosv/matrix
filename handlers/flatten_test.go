package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestFlatten_WithValidMatrix_ShouldOutputExpected(t *testing.T) {
	body := "1,2,3\n4,5,6\n7,8,9"

	request := createRequest(t, body)
	response := httptest.NewRecorder()

	FlattenHandler(response, request)

	responseBody := response.Body.String()

	assert.Equal(t, "1,2,3,4,5,6,7,8,9", responseBody)
}

func TestFlatten_WithNonNumbers_ShouldOutputError(t *testing.T) {
	body := "a"

	request := createRequest(t, body)
	response := httptest.NewRecorder()

	FlattenHandler(response, request)

	responseBody := response.Body.String()

	assert.Equal(t, "error all values should be numbers", responseBody)
}

func TestFlatten_WithNonCsv_ShouldOutputError(t *testing.T) {
	body := "1,2,3\n4,5,6\n7,8"

	request := createRequest(t, body)
	response := httptest.NewRecorder()

	FlattenHandler(response, request)

	responseBody := response.Body.String()

	assert.Equal(t, "error could not read the file", responseBody)
}

func TestFlatten_WithNonMatrix_ShouldOutputError(t *testing.T) {
	body := "1,2,3\n4,5,6\n7,8,9\n10,11,12"

	request := createRequest(t, body)
	response := httptest.NewRecorder()

	FlattenHandler(response, request)

	responseBody := response.Body.String()

	assert.Equal(t, "error the number of columns and rows should be the same", responseBody)
}
