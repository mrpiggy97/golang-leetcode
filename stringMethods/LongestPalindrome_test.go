package stringMethods_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/mrpiggy97/golang-leetcode/stringMethods"
)

func testIsPalindrome(testCase *testing.T) {
	var testingStr string = "aba"
	var result bool = stringMethods.IsPalindrome(strings.Split(testingStr, ""))
	var expectedResult bool = true
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	testingStr = "aas43"
	result = stringMethods.IsPalindrome(strings.Split(testingStr, ""))
	expectedResult = false
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	testingStr = "anitalavalatina"
	result = stringMethods.IsPalindrome(strings.Split(testingStr, ""))
	expectedResult = true
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
}

func testGetClosureIndexes(testCase *testing.T) {
	var testingStr string = "fabi1011f"
	var testingSlice []string = strings.Split(testingStr, "")
	var result []int = stringMethods.GetClosureIndexes(0, strings.Split(testingStr, ""))
	if len(result) != 1 {
		testCase.Errorf("expected length of result to be 1, got %d, instead", len(result))
		os.Exit(2)
	}
	if result[0] != len(testingSlice)-1 {
		testCase.Errorf("expected index to be %d got %d instead", len(testingSlice)-1, result[0])
	}

	result = stringMethods.GetClosureIndexes(2, testingSlice)
	if len(result) != 0 {
		testCase.Errorf("expected length of result to be 1, got %d, instead", len(result))
		os.Exit(2)
	}
}

func testGetSubString(testCase *testing.T) {
	var testingString string = "testingsubstr"
	var result []string = stringMethods.GetSubStr(strings.Split(testingString, ""), 1, 4)
	if len(result) != 4 {
		fmt.Println(result)
		testCase.Errorf("expected length of result to be 4, got %d instead", len(result))
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
	testCase.Run("action=test-get-closure-indexes", testGetClosureIndexes)
	testCase.Run("action=test-get-sub-string", testGetSubString)
	testCase.Run("action=test-longest-palindrome", testLongestPalindrome)
}
