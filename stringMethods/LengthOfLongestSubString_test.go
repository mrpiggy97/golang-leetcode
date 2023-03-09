package stringMethods_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/stringMethods"
)

func testLengthOfLongestSubString(testCase *testing.T) {
	var testStr string = "fenoded"
	var expectedResult int = 5
	var result int = stringMethods.LengthOfLongestSubString(testStr)
	if result != expectedResult {
		testCase.Errorf("expected %d to be %d", result, expectedResult)
	}

	testStr = "dvdf"
	expectedResult = 3
	result = stringMethods.LengthOfLongestSubString(testStr)
	if result != expectedResult {
		testCase.Errorf("expected %d to be %d", result, expectedResult)
	}

	testStr = "refofem"
	expectedResult = 4
	result = stringMethods.LengthOfLongestSubString(testStr)
	if result != expectedResult {
		testCase.Errorf("expected %d to be %d", result, expectedResult)
	}
}

func TestLengthOfLongestSubString(testCase *testing.T) {
	testCase.Run("action=test-longest-substring", testLengthOfLongestSubString)
}
