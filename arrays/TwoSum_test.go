package arrays_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/arrays"
)

func testTwoSum(testCase *testing.T) {
	var testArray1 []int = []int{2, 7, 11, 15}
	var testTarget1 int = 9
	var expectedResult []int = []int{0, 1}
	var result []int = arrays.TwoSum(testArray1, testTarget1)
	if expectedResult[0] != result[0] {
		testCase.Errorf("expected first index to be %d, got %d instead", expectedResult[0], result[0])
	}
	if expectedResult[1] != result[1] {
		testCase.Errorf("expected first index to be %d, got %d instead", expectedResult[1], result[1])
	}
}

func TestTwoSum(testCase *testing.T) {
	testCase.Run("action=test-two-sum", testTwoSum)
}
