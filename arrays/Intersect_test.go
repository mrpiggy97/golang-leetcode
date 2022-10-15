package arrays_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/arrays"
)

func testIntersection(testCase *testing.T) {
	var nums1 []int = []int{1, 8, 7, 9, 2, 9, 100, 9, 9, 9}
	var nums2 []int = []int{1, 9, 33, 5, 190, 9, 9, 177, 12, 122}
	var result []int = arrays.Intersect(nums1, nums2)
	var expectedResult map[int]int = make(map[int]int)
	expectedResult[1] = 1
	expectedResult[9] = 3
	if len(result) != 4 {
		testCase.Errorf("length of slices is not the same")
	}
	for _, val := range result {
		count, exists := expectedResult[val]
		if !exists {
			testCase.Errorf("expected %d to exist in expected results", val)
		}
		if exists {
			if val == 9 {
				if count != 3 {
					testCase.Errorf("expected count for %d to be %d, got %d instead", val, 3, count)
				}
			}
			if val == 1 {
				if count != 1 {
					testCase.Errorf("expected count for %d to be %d, got %d instead", val, 1, count)
				}
			}
		}
	}
}

func TestIntersection(testCase *testing.T) {
	testCase.Run("action=test-intersection", testIntersection)
}
