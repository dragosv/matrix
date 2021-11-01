package processing

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"
)

type Matrix struct {
}

type MatrixReader interface {
	Read(file io.Reader) ([][]int, error)
}

//Read returns the file as an integer matrix.
func (_ *Matrix) Read(file io.Reader) ([][]int, error) {
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, errors.New("could not read the file")
	}

	var intRecords [][]int

	intRecords = make([][]int, len(records))

	for i := range records {
		intRecords[i] = make([]int, len(records))
	}

	for i, record := range records {
		if len(record) != len(records) {
			return nil, errors.New("the number of columns and rows should be the same")
		}

		for j, value := range record {
			intValue, err := strconv.Atoi(value)

			if err != nil {
				return nil, errors.New("all values should be numbers")
			}

			intRecords[i][j] = intValue
		}
	}

	return intRecords, nil
}
