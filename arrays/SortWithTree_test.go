package arrays_test

import (
	"fmt"
	"testing"

	"github.com/mrpiggy97/golang-leetcode/arrays"
)

func testSortWithTree(testCase *testing.T) {
	var testingSlice []int = []int{-1, 2, 77, 20, -1, 15}
	var expectedResult []int = []int{-1, 2, 15, 20, -1, 77}
	var result = arrays.SortWithTree(testingSlice)
	fmt.Println(result)
	for index, value := range result {
		if value != expectedResult[index] {
			testCase.Errorf("expected %d to be %d", value, expectedResult[index])
		}
	}
}

func TestSortWithTree(testCase *testing.T) {
	testCase.Run("action=test-sort-with-tree", testSortWithTree)
}
