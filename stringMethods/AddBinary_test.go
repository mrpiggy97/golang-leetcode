package stringMethods_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/stringMethods"
)

func testAddBinary(testCase *testing.T) {
	var a string = "100"
	var b string = "11"
	var result string = stringMethods.AddBinary(a, b)
	var expectedResult string = "111"
	if result != expectedResult {
		testCase.Errorf("expected %s to be %s", result, expectedResult)
	}
	a = "111"
	b = "11"
	expectedResult = "1010"
	result = stringMethods.AddBinary(a, b)
	if result != expectedResult {
		testCase.Errorf("expected %s to be %s", result, expectedResult)
	}
}

func TestAddBinary(testCase *testing.T) {
	testCase.Run("action=test-add-binary", testAddBinary)
}
