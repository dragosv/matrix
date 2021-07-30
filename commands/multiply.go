package commands

import (
	"errors"
	"github.com/JohnCGriffin/overflow"
	"strconv"
)

func Multiply(records [][]int) (string, error) {
	var total int

	if len(records) > 0 {
		total = 1
	}

	var ok bool
	for _, record := range records {
		for _, value := range record {
			total, ok = overflow.Mul(total, value)

			if !ok {
				return "", errors.New("multiply overflow")
			}
		}
	}

	return strconv.Itoa(total), nil
}
