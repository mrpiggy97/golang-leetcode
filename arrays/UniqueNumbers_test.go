package arrays_test

import (
	"fmt"
	"testing"

	"github.com/mrpiggy97/golang-leetcode/arrays"
)

func testUniqueNumbers(testCase *testing.T) {
	var numbers []int = []int{1, 2, 1, 2, -1, 45, 6, 71, 6}
	var result []int = arrays.UniqueNumbers(numbers)
	var expectedResult []int = []int{-1, 45, 71}
	if len(result) != len(expectedResult) {
		fmt.Println(result)
		testCase.Errorf("expected length of result to be %d, got %d, instead", len(result), len(expectedResult))
	}
}

func TestUniqueNumbers(testCase *testing.T) {
	testCase.Run("action=test-unique-numbers", testUniqueNumbers)
}
