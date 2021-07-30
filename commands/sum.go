package commands

import (
	"errors"
	"github.com/JohnCGriffin/overflow"
	"strconv"
)

func Sum(records [][]int) (string, error) {
	var total int

	var ok bool
	for _, record := range records {
		for _, value := range record {
			total, ok = overflow.Add(total, value)

			if !ok {
				return "", errors.New("sum overflow")
			}
		}
	}

	return strconv.Itoa(total), nil
}
