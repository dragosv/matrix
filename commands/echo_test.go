package commands

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEcho_Empty_ShouldOutputExpected(t *testing.T) {
	var records [][]int

	matrix := Matrix{}

	outputText, err := matrix.Echo(records)

	assert.Nil(t, err)
	assert.Equal(t, "", outputText)
}

func TestEcho_NonEmpty_ShouldOutputExpected(t *testing.T) {
	records := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	matrix := Matrix{}

	outputText, err := matrix.Echo(records)

	assert.Nil(t, err)
	assert.Equal(t, "1,2,3,4\n5,6,7,8\n9,10,11,12\n13,14,15,16", outputText)
}
