package commands

import (
	"strconv"
	"strings"
)

//Flatten returns the matrix as a 1 line string, with values separated by commas.
func Flatten(matrix [][]int) (string, error) {
	var lines []string

	lines = make([]string, len(matrix))

	for i, row := range matrix {
		var stringValues []string

		stringValues = make([]string, len(row))

		for j, value := range row {
			stringValue := strconv.Itoa(value)
			stringValues[j] = stringValue
		}

		lines[i] = strings.Join(stringValues, ",")
	}

	return strings.Join(lines, ","), nil
}
