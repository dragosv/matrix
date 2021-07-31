package commands

import (
	"errors"
	"github.com/JohnCGriffin/overflow"
	"strconv"
)

//Multiply returns the product of the integers in the matrix.
//Will return multiply overflow if the multiplication overflows.
func Multiply(matrix [][]int) (string, error) {
	var total int

	if len(matrix) > 0 {
		total = 1
	}

	var ok bool
	for _, row := range matrix {
		for _, value := range row {
			total, ok = overflow.Mul(total, value)

			if !ok {
				return "", errors.New("multiply overflow")
			}
		}
	}

	return strconv.Itoa(total), nil
}
