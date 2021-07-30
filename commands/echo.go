package commands

import (
	"bytes"
	"strconv"
	"strings"
)

func Echo(records [][]int) (string, error) {
	var buffer bytes.Buffer

	for i, record := range records {
		if i > 0 {
			buffer.WriteString("\n")
		}

		var stringValues []string

		stringValues = make([]string, len(record))

		for j, value := range record {
			stringValue := strconv.Itoa(value)
			stringValues[j] = stringValue
		}

		buffer.WriteString(strings.Join(stringValues, ","))
	}

	return buffer.String(), nil
}
