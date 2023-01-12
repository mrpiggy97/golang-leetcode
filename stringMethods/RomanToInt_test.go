package stringMethods_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/stringMethods"
)

func testExpectedSum(testCase *testing.T) {
	var testing_number string = "IV"
	var expectedResult int = 4
	var result int = stringMethods.RomanToInt(testing_number)
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
}

func TestRomanToInt(testCase *testing.T) {
	testCase.Run("action=test-expected-sum", testExpectedSum)
}
