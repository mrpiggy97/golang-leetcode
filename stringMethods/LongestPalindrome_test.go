package stringMethods_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/stringMethods"
)

func testIsPalindrome(testCase *testing.T) {
	var testingStr string = "aba"
	var result bool = stringMethods.IsPalindrome(testingStr, 0, len(testingStr)-1)
	var expectedResult bool = true
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	testingStr = "aas43"
	result = stringMethods.IsPalindrome(testingStr, 3, len(testingStr)-1)
	expectedResult = false
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	testingStr = "anitalavalatina"
	result = stringMethods.IsPalindrome(testingStr, 0, len(testingStr)-1)
	expectedResult = true
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	testingStr = "abafafaaaaaf"
	result = stringMethods.IsPalindrome(testingStr, 5, 11)
	//fmt.Println(string(testingStr[5:11] + string(testingStr[11])))
	expectedResult = true
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
}

func testLongestPalindrome(testCase *testing.T) {
	var testingString string = "abafafaaaaaf"
	var result string = stringMethods.LongestPalindrome(testingString)
	var expectedResult = "faaaaaf"
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	testingString = "colac"
	result = stringMethods.LongestPalindrome(testingString)
	expectedResult = "c"
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	testingString = "a"
	result = stringMethods.LongestPalindrome(testingString)
	expectedResult = "a"
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
}

func TestLongestPalindrome(testCase *testing.T) {
	testCase.Run("action=test-is-palindrome", testIsPalindrome)
	testCase.Run("action=test-longest-palindrome", testLongestPalindrome)
}
