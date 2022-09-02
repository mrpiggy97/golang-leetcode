package arrays_test

import (
	"fmt"
	"testing"

	"github.com/mrpiggy97/golang-leetcode/arrays"
)

func testRemoveItemFromSlice(testCase *testing.T) {
	var testingSlice []int = []int{3, 55, 3, 66}
	testingSlice = arrays.RemoveItemFromSlice(2, testingSlice)
	if len(testingSlice) != 3 {
		testCase.Errorf("expected length to be %d, not %d", 3, len(testingSlice))
	}
}

func testSortSlice(testCase *testing.T) {
	var testingArray []int = []int{55, 3, 4, 1, 0, 4}
	var expectedResult []int = []int{0, 1, 3, 4, 4, 55}
	testingArray = arrays.SortSlice(testingArray)
	for index, value := range testingArray {
		if testingArray[index] != expectedResult[index] {
			testCase.Errorf("expected %d to be %d", value, expectedResult[index])
		}
	}
}

func testMerge(testCase *testing.T) {
	var testArray1 []int = []int{3, 5, 1, 66, 0, 0}
	var testArray2 []int = []int{7, 3}
	var expectedResult []int = []int{1, 3, 3, 5, 7, 66}
	arrays.Merge(testArray1, len(testArray1), testArray2, len(testArray2))
	fmt.Println(testArray1)
	for index, value := range testArray1 {
		if value != expectedResult[index] {
			testCase.Errorf("expected %d to be %d", value, expectedResult[index])
		}
	}
}

func TestMerge(testCase *testing.T) {
	testCase.Run("action=test-remove-item-from-slice", testRemoveItemFromSlice)
	testCase.Run("action=test-sort-slice", testSortSlice)
	testCase.Run("action=test-merge", testMerge)
}
