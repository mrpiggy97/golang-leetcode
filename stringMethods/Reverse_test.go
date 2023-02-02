package stringMethods_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/stringMethods"
)

func testReverse(testCase *testing.T) {
	var testingString string = "foo(bar)ram"
	var expectedResult string = "foorabram"
	var result string = stringMethods.FilterStringsToReverse(testingString)
	if result != expectedResult {
		testCase.Errorf("expected %s to be %s", result, expectedResult)
	}

	testingString = "(bar)"
	expectedResult = "rab"
	result = stringMethods.FilterStringsToReverse(testingString)
	if result != expectedResult {
		testCase.Errorf("expected %s to be %s", result, expectedResult)
	}

	testingString = "(u(love)i)"
	expectedResult = "iloveu"
	result = stringMethods.FilterStringsToReverse(testingString)
	if result != expectedResult {
		testCase.Errorf("expected %s to be %s", result, expectedResult)
	}
}

func TestReverse(testCase *testing.T) {
	testCase.Run("action=test-reverse", testReverse)
}
