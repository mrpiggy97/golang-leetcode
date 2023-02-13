package stringMethods_test

import (
	"fmt"
	"testing"

	"github.com/mrpiggy97/golang-leetcode/stringMethods"
)

func testGetStart(testCase *testing.T) {
	var testStr string = "BCA"
	var expectedStart int = 703
	var expectedEnd int = 18279
	var expectedFinalString string = "AAA"
	start, end, finalString := stringMethods.GetStart(testStr)
	if start != expectedStart {
		fmt.Println(start, end, finalString)
		testCase.Errorf("expected %d to be %d", start, expectedStart)
	}

	if end != expectedEnd {
		testCase.Errorf("expected %d to be %d", end, expectedEnd)
	}

	if finalString != expectedFinalString {
		testCase.Errorf("expected %s to be %s", finalString, expectedFinalString)
	}
}

func testTitleToNumber(testCase *testing.T) {
	var testStr string = "AAA"
	var expectedResult int = 703
	var result int = stringMethods.TitleToNumber(testStr)
	if result != expectedResult {
		testCase.Errorf("expected %d to be %d", result, expectedResult)
	}

	testStr = "BCA"
	expectedResult = 1431
	result = stringMethods.TitleToNumber(testStr)
	if result != expectedResult {
		testCase.Errorf("expected %d to be %d", result, expectedResult)
	}

	testStr = "DFI"
	expectedResult = 2869
	result = stringMethods.TitleToNumber(testStr)
	if result != expectedResult {
		testCase.Errorf("expected %d to be %d", result, expectedResult)
	}
}

func TestTitleToNumber(testCase *testing.T) {
	testCase.Run("action=test-title-to-number", testTitleToNumber)
	testCase.Run("action=test-get-start", testGetStart)
}
