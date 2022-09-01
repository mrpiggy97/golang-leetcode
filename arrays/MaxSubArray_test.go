package arrays_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/arrays"
)

func testMaxSubArray(testCase *testing.T) {
	var array []int = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	var secondSlice []int = []int{1}
	var thirdSlice []int = []int{-1}
	var fourthSlice []int = []int{-2, -1}
	var highestSum int = arrays.MaxSubArray(array)
	var expectedResult int = 6
	if highestSum != expectedResult {
		testCase.Errorf("expected %d, to be %d", highestSum, expectedResult)
	}
	highestSum = arrays.MaxSubArray(secondSlice)
	expectedResult = 1
	if highestSum != expectedResult {
		testCase.Errorf("expected %d, to be %d", highestSum, expectedResult)
	}
	highestSum = arrays.MaxSubArray(thirdSlice)
	expectedResult = -1

	if highestSum != expectedResult {
		testCase.Errorf("expected %d, to be %d", highestSum, expectedResult)
	}

	highestSum = arrays.MaxSubArray(fourthSlice)
	expectedResult = -1
	if highestSum != expectedResult {
		testCase.Errorf("expected %d, to be %d", highestSum, expectedResult)
	}
}

func TestMaxSubArray(testCase *testing.T) {
	testCase.Run("action=test-max-sub-array", testMaxSubArray)
}
