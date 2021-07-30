package commands

import (
	"strconv"
	"strings"
)

func Flatten(records [][]int) (string, error) {
	var lines []string

	lines = make([]string, len(records))

	for i, record := range records {
		var stringValues []string

		stringValues = make([]string, len(record))

		for j, value := range record {
			stringValue := strconv.Itoa(value)
			stringValues[j] = stringValue
		}

		lines[i] = strings.Join(stringValues, ",")
	}

	return strings.Join(lines, ","), nil
}
