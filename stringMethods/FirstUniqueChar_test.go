package stringMethods_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/stringMethods"
)

func testFirstUniqueChar(testCase *testing.T) {
	var testString string = "aabb"
	var result int = stringMethods.FirstUniqueChar(testString)
	var expectedResult int = -1
	if result != expectedResult {
		testCase.Errorf("expected %d to be %d", result, expectedResult)
	}
	testString = "leetcode"
	expectedResult = 0
	result = stringMethods.FirstUniqueChar(testString)
	if result != expectedResult {
		testCase.Errorf("expected %d to be %d", result, expectedResult)
	}
	testString = "loveleetcode"
	expectedResult = 2
	result = stringMethods.FirstUniqueChar(testString)
	if result != expectedResult {
		testCase.Errorf("expected %d to be %d", result, expectedResult)
	}
}

func TestFirstUniqueChar(testCase *testing.T) {
	testCase.Run("action=test-first-unique-char", testFirstUniqueChar)
}
