package stringMethods_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/stringMethods"
)

func testIsPalindromeLeetcode(testCase *testing.T) {
	var testingString string = "!faf"
	var expectedResult bool = true
	var result bool = stringMethods.IsPalindromeLeetcode(testingString)
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	testingString = "!!!??"
	expectedResult = false
	result = stringMethods.IsPalindromeLeetcode(testingString)

	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	testingString = "!f  f!"
	expectedResult = true
	result = stringMethods.IsPalindromeLeetcode(testingString)
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
}

func TestIsPalindrome(testCase *testing.T) {
	testCase.Run("action=test-is-palindrome-leetcode", testIsPalindromeLeetcode)
}
