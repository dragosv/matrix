package processing

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestMatrix_WithValidData_ShouldOutputExpected(t *testing.T) {
	body := "1,2\n3,4"

	reader := createReader(body)

	matrix, err := Matrix(reader)
	assert.Nil(t, err)

	assert.Equal(t, 1, matrix[0][0])
	assert.Equal(t, 2, matrix[0][1])
	assert.Equal(t, 3, matrix[1][0])
	assert.Equal(t, 4, matrix[1][1])
}

func TestMatrix_WithNonNumbers_ShouldOutputError(t *testing.T) {
	body := "a"

	reader := createReader(body)

	_, err := Matrix(reader)
	assert.NotNil(t, err)
	assert.Equal(t, "all values should be numbers", err.Error())
}

func TestMatrix_WithNonCsv_ShouldOutputError(t *testing.T) {
	body := "1,2,3\n4,5,6\n7,8"

	reader := createReader(body)

	_, err := Matrix(reader)
	assert.NotNil(t, err)
	assert.Equal(t, "could not read the file", err.Error())
}

func TestMatrix_WithNonMatrix_ShouldOutputError(t *testing.T) {
	body := "1,2,3\n4,5,6\n7,8,9\n10,11,12"

	reader := createReader(body)

	_, err := Matrix(reader)
	assert.NotNil(t, err)
	assert.Equal(t, "the number of columns and rows should be the same", err.Error())
}

func createReader(body string) io.Reader {
	bodyBytes := []byte(body)

	return bytes.NewReader(bodyBytes)
}
