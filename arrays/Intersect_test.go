package arrays_test

import (
	"fmt"
	"testing"

	"github.com/mrpiggy97/golang-leetcode/arrays"
)

func testGetElementsInCommon(testCase *testing.T) {
	var testSlice1 []int = []int{222, 1, 6, 7, 10}
	var testSlice2 []int = []int{4, 6, 6, 6, 91, 100, 222, 1}
	var result []int = arrays.GetElementsInCommon[int](testSlice1, testSlice2)
	var expectedResult []int = []int{222, 1, 6}
	for index, value := range result {
		if value != expectedResult[index] {
			testCase.Errorf("expected %d to be %d", value, expectedResult[index])
		}
	}
}

func testIntersection(testCase *testing.T) {
	var nums1 []int = []int{1, 8, 7, 9, 2, 9, 100, 9, 9, 9}
	var nums2 []int = []int{1, 9, 33, 5, 190, 9, 9, 177, 12, 122}
	var expectedResult = []int{1, 9, 9, 9}
	var result []int = arrays.Intersect(nums1, nums2)
	if len(result) != len(expectedResult) {
		testCase.Errorf("length of slices is not the same")
	}
	for index, value := range expectedResult {
		if value != result[index] {
			testCase.Errorf("expected %d to be %d", value, expectedResult[index])
		}
	}
	fmt.Println(result)
}

func TestIntersection(testCase *testing.T) {
	testCase.Run("action=test-get-elements-in-common", testGetElementsInCommon)
	testCase.Run("action=test-intersection", testIntersection)
}
