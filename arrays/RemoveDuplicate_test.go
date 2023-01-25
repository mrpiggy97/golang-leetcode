package arrays_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/arrays"
)

func testRemoveDuplicate(testCase *testing.T) {
	var testSlice []int = []int{1, 1, 2, 2, 3, 4}
	var expectedResult int = 4
	var result int = arrays.RemoveDuplicates(testSlice)
	if expectedResult != result {
		testCase.Errorf("expected %d to be %d", result, expectedResult)
	}
}

func TestRemoveDuplicate(testCase *testing.T) {
	testCase.Run("action=test-remove-duplicate", testRemoveDuplicate)
}
