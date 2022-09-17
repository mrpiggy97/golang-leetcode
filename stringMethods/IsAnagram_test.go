package stringMethods_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/stringMethods"
)

func testIsAnagram(testCase *testing.T) {
	var testingAnagram string = "anagram"
	var testingSecondString string = "nagaram"
	var result bool = stringMethods.IsAnagram(testingAnagram, testingSecondString)
	var expectedResult bool = true

	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	testingSecondString = "nagaramd"
	result = stringMethods.IsAnagram(testingAnagram, testingSecondString)
	expectedResult = false
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	testingSecondString = "zzzzzzzz"
	result = stringMethods.IsAnagram(testingAnagram, testingSecondString)
	expectedResult = false
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
}

func TestIsAnagram(testCase *testing.T) {
	testCase.Run("action=test-is-anagram", testIsAnagram)
}
