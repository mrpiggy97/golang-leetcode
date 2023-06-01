package arrays_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/arrays"
)

func testGetMinCost(testCase *testing.T) {
	var testSlice []int = []int{2, 2, 1, 0}
	var result int = arrays.GetMinCost(testSlice, 1)
	var expectedResult int = 2
	if result != expectedResult {
		testCase.Errorf("expected %d to be %d", result, expectedResult)
	}
}

func testMinCostClimbingStairs(testCase *testing.T) {
	var cost []int = []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	var result int = arrays.MinCostClimbingStairs(cost)
	var expectedResult int = 6
	if result != expectedResult {
		testCase.Errorf("expected %d to be %d", result, expectedResult)
	}
}

func TestMinCostClimbingStairs(testCase *testing.T) {
	testCase.Run("action=test-get-min-cost", testGetMinCost)
	testCase.Run("action=test-mincost-climbing-stairs", testMinCostClimbingStairs)
}
