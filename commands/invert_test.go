package commands

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvert_Empty_ShouldOutputExpected(t *testing.T) {
	var records [][]int

	matrix := Matrix{}

	outputText, err := matrix.Invert(records)

	assert.Nil(t, err)
	assert.Equal(t, "", outputText)
}

func TestInvert_NonEmpty_ShouldOutputExpected(t *testing.T) {
	records := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	matrix := Matrix{}

	outputText, err := matrix.Invert(records)

	assert.Nil(t, err)
	assert.Equal(t, "1,5,9,13\n2,6,10,14\n3,7,11,15\n4,8,12,16", outputText)
}
