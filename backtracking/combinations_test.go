package backtracking_test

import (
	"testing"

	"github.com/jerensl/algorithm/backtracking"
	"github.com/stretchr/testify/assert"
)

func TestCombinations(t *testing.T) {
	expected := [][]int{
		{1,2},
		{1,3},
		{1,4},
		{2,3},
		{2,4},
		{3,4},
	}
	result := backtracking.Combinations(4, 2)
	assert.Equal(t, expected, result, "Return value of 2D array should be %v, but got %v", expected, result)
}