package commands

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestMultiply_Empty_ShouldOutputExpected(t *testing.T) {
	var records [][]int

	matrix := Matrix{}

	outputText, err := matrix.Multiply(records)

	assert.Nil(t, err)
	assert.Equal(t, "0", outputText)
}

func TestMultiply_NonEmpty_ShouldOutputExpected(t *testing.T) {
	records := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	matrix := Matrix{}

	outputText, err := matrix.Multiply(records)

	assert.Nil(t, err)
	assert.Equal(t, "20922789888000", outputText)
}

func TestMultiply_AboveMaxInt_ShouldOverflow(t *testing.T) {
	records := [][]int{
		{math.MaxInt64, 2},
		{1, 1},
	}

	matrix := Matrix{}

	_, err := matrix.Multiply(records)

	assert.NotNil(t, err)
	assert.Equal(t, "multiply overflow", err.Error())
}
