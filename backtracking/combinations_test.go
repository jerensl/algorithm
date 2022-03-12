package backtracking_test

import (
	"testing"

	"github.com/jerensl/algorithm/backtracking"
	"github.com/stretchr/testify/assert"
)

func TestCombinations(t *testing.T) {
	result := backtracking.Combinations(4, 2)
	assert.NotNil(t, result, "Return value of 2D array should not be nil, but got %v", result)
}