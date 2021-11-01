package commands

import (
	"errors"
	"github.com/JohnCGriffin/overflow"
	"strconv"
)

//Sum returns the sum of the integers in the matrix.
//Will return sum overflow error if the addition overflows.
func (_ *Matrix) Sum(matrix [][]int) (string, error) {
	var total int

	var ok bool
	for _, row := range matrix {
		for _, value := range row {
			total, ok = overflow.Add(total, value)

			if !ok {
				return "", errors.New("sum overflow")
			}
		}
	}

	return strconv.Itoa(total), nil
}
