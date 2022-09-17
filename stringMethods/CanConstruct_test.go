package stringMethods_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/stringMethods"
)

func testCanConstruct(testCase *testing.T) {
	var testNote string = "abcc"
	var testMagazine string = "abbc"
	var result bool = stringMethods.CanConstruct(testNote, testMagazine)
	var expectedResult bool = false
	if result != expectedResult {
		testCase.Errorf("expected %v %v", result, expectedResult)
	}
	testNote = "aa"
	testMagazine = "aab"
	result = stringMethods.CanConstruct(testNote, testMagazine)
	expectedResult = true
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
	testNote = "ab"
	testMagazine = "zmk"
	result = stringMethods.CanConstruct(testNote, testMagazine)
	expectedResult = false
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
	testNote = "bg"
	testMagazine = "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj"
	result = stringMethods.CanConstruct(testNote, testMagazine)
	expectedResult = true
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
}

func TestCanConstruct(testCase *testing.T) {
	testCase.Run("action=test-can-construct", testCanConstruct)
}
