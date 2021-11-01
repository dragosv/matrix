package commands

import (
	"bytes"
	"strconv"
	"strings"
)

//Echo returns the matrix as a string in matrix format.
func (_ *Matrix) Echo(matrix [][]int) (string, error) {
	var buffer bytes.Buffer

	for i, row := range matrix {
		if i > 0 {
			buffer.WriteString("\n")
		}

		var stringValues []string

		stringValues = make([]string, len(row))

		for j, value := range row {
			stringValue := strconv.Itoa(value)
			stringValues[j] = stringValue
		}

		buffer.WriteString(strings.Join(stringValues, ","))
	}

	return buffer.String(), nil
}
