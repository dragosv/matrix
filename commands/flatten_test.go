package commands

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlatten_Empty_ShouldOutputExpected(t *testing.T) {
	var records [][]int

	matrix := Matrix{}

	outputText, err := matrix.Flatten(records)

	assert.Nil(t, err)
	assert.Equal(t, "", outputText)
}

func TestFlatten_NonEmpty_ShouldOutputExpected(t *testing.T) {
	records := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	matrix := Matrix{}

	outputText, err := matrix.Flatten(records)

	assert.Nil(t, err)
	assert.Equal(t, "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16", outputText)
}
