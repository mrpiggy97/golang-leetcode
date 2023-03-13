package linkedlists_test

import (
	"testing"

	linkedlists "github.com/mrpiggy97/golang-leetcode/linkedLists"
)

func testSumNumbers(testCase *testing.T) {
	var num1 string = "9999"
	var num2 string = "99"
	var expectedResult string = "10098"
	var result string = linkedlists.SumNumbers(num1, num2)
	if result != expectedResult {
		testCase.Errorf("expected %s to be %s", result, expectedResult)
	}

	num1 = "543"
	num2 = "666"
	expectedResult = "1209"
	result = linkedlists.SumNumbers(num1, num2)
	if result != expectedResult {
		testCase.Errorf("expected %s to be %s", result, expectedResult)
	}

	num1 = "3"
	num2 = "15"
	expectedResult = "18"
	result = linkedlists.SumNumbers(num1, num2)
	if result != expectedResult {
		testCase.Errorf("expected %s to be %s", result, expectedResult)
	}
}

func TestTwoNumbers(testCase *testing.T) {
	testCase.Run("Action=test-sum-numbers", testSumNumbers)
}
