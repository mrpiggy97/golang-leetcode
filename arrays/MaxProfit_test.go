package arrays_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/arrays"
)

func testMaxProfit(testCase *testing.T) {

	testingArray := []int{4, 7, 1, 2}
	bestProfit := arrays.MaxProfit(testingArray)
	expectedResult := 3
	if bestProfit != expectedResult {
		testCase.Errorf("expected %d to be %d", bestProfit, expectedResult)
	}

	testingArray = []int{7, 1, 5, 3, 6, 4}
	bestProfit = arrays.MaxProfit(testingArray)
	expectedResult = 5
	if bestProfit != expectedResult {
		testCase.Errorf("expected %d to be %d", bestProfit, expectedResult)
	}

	testingArray = []int{7, 6, 4, 3, 1}
	bestProfit = arrays.MaxProfit(testingArray)
	expectedResult = 0
	if bestProfit != expectedResult {
		testCase.Errorf("expected %d to be %d", bestProfit, expectedResult)
	}

	testingArray = []int{4, 11, 1, 2, 7}
	bestProfit = arrays.MaxProfit(testingArray)
	expectedResult = 7
	if bestProfit != expectedResult {
		testCase.Errorf("expected %d to be %d", bestProfit, expectedResult)
	}
}

func TestMaxProfit(testCase *testing.T) {
	testCase.Run("action=test-max-profit", testMaxProfit)
}
