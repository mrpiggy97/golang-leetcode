package arrays_test

import (
	"fmt"
	"testing"

	"github.com/mrpiggy97/golang-leetcode/arrays"
)

func testGetNewLevel(testCase *testing.T) {
	var testingSlice []int = []int{1, 1}
	var newLevel []int = arrays.GetNewLevel(testingSlice)
	var expectedLength int = 3
	if len(newLevel) != expectedLength {
		testCase.Errorf("expected length of array to be %d, got %d instead", expectedLength, len(newLevel))
	}
	newLevel = arrays.GetNewLevel(newLevel)
	expectedLength = 4
	if len(newLevel) != expectedLength {
		testCase.Errorf("expected length of array to be %d, got %d instead", expectedLength, len(newLevel))
	}
}

func testGenerate(testCase *testing.T) {
	var levels int = 3
	var pascalLevels [][]int = arrays.Generate(levels)
	var expectedLength int = 3
	if len(pascalLevels) != expectedLength {
		testCase.Errorf("expected length of pascal levels to be %d not %d", expectedLength, len(pascalLevels))
	}
	levels = 6
	pascalLevels = arrays.Generate(levels)
	expectedLength = 6
	if len(pascalLevels) != expectedLength {
		testCase.Errorf("expected length of pascal levels to be %d not %d", expectedLength, len(pascalLevels))
	}
	fmt.Println(pascalLevels)
}

func TestGenerate(testCase *testing.T) {
	testCase.Run("action=test-get-new-level", testGetNewLevel)
	testCase.Run("action=test-generate", testGenerate)
}
