package arrays_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/arrays"
)

func testRemoveElement(testCase *testing.T) {
	var testSlice []int = []int{3, 3, 4, 5, 6}
	var expectedResult int = 3
	var result int = arrays.RemoveElement(testSlice, 3)
	if result != expectedResult {
		testCase.Errorf("expected %d to be %d", result, expectedResult)
	}
}

func TestRemoveElement(testCase *testing.T) {
	testCase.Run("action=test-remove-element", testRemoveElement)
}
