package commands

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestSum_Empty_ShouldOutputExpected(t *testing.T) {
	var records [][]int

	matrix := Matrix{}

	outputText, err := matrix.Sum(records)

	assert.Nil(t, err)
	assert.Equal(t, "0", outputText)
}

func TestSum_Standard_ShouldOutputExpected(t *testing.T) {
	records := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	matrix := Matrix{}

	outputText, err := matrix.Sum(records)

	assert.Nil(t, err)
	assert.Equal(t, "136", outputText)
}

func TestSum_AboveMaxInt_ShouldOverflow(t *testing.T) {
	records := [][]int{
		{math.MaxInt64, 1},
		{1, 1},
	}

	matrix := Matrix{}

	_, err := matrix.Sum(records)

	assert.NotNil(t, err)
	assert.Equal(t, "sum overflow", err.Error())
}
