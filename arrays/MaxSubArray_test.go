package arrays_test

import (
	"fmt"
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

func testGetSliceWithHighestSum(testCase *testing.T) {
	var firstSlice []int = []int{-1, 3, 4, 1}
	var secondSlice []int = []int{-1, 3, 4}
	var slices [][]int = [][]int{firstSlice, secondSlice}
	_, sum := arrays.GetSliceWithHighestSum(slices)
	var expectedResult int = 7
	if sum != expectedResult {
		testCase.Errorf("expected %d to be %d", sum, expectedResult)
	}
}

func testGetSlicesFromIndex(testCase *testing.T) {
	var testingSlice []int = []int{2, -1, 44, 5, 6}
	var slices [][]int = arrays.GetSlicesFromIndex(1, testingSlice)
	var length int = 4
	fmt.Println(slices)
	if len(slices) != length {
		testCase.Errorf("length is %d, expected it to be %d", len(slices), length)
	}
}

func TestMaxSubArray(testCase *testing.T) {
	testCase.Run("action=test-max-sub-array", testMaxSubArray)
	testCase.Run("action=test-get-max-sum", testGetSliceWithHighestSum)
	testCase.Run("action=test-get-slices-from-index", testGetSlicesFromIndex)
}
