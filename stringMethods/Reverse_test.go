package stringMethods_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/stringMethods"
)

func testReverse(testCase *testing.T) {
	var testingString string = "foo(bar(baz))ram"
	var expectedResult string = "foobazrabram"
	var result string = stringMethods.FilterStringsToReverse(testingString)
	if result != expectedResult {
		testCase.Errorf("expected %s to be %s", result, expectedResult)
	}

	testingString = "(foo)zam"
	expectedResult = "oofzam"
	result = stringMethods.FilterStringsToReverse(testingString)
	if result != expectedResult {
		testCase.Errorf("expected %s to be %s", result, expectedResult)
	}

	testingString = "zam(foo)"
	expectedResult = "zamoof"
	result = stringMethods.FilterStringsToReverse(testingString)
	if result != expectedResult {
		testCase.Errorf("expected %s to be %s", result, expectedResult)
	}

	testingString = "()"
	expectedResult = ""
	result = stringMethods.FilterStringsToReverse(testingString)
	if result != expectedResult {
		testCase.Errorf("expected %s to be %s", result, expectedResult)
	}

	testingString = "(foo)(ram)(zam)"
	expectedResult = "oofmarmaz"
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
