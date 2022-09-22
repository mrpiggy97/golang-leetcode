package queue_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/queue"
)

func testIsOpener(testCase *testing.T) {
	var testingValue string = "{"
	var result bool = queue.IsOpener(rune(testingValue[0]))
	var expectedResult bool = true
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	testingValue = "("
	result = queue.IsOpener(rune(testingValue[0]))
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	testingValue = "["
	result = queue.IsOpener(rune(testingValue[0]))
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	testingValue = "}"
	expectedResult = false
	result = queue.IsOpener(rune(testingValue[0]))
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
}

func testIsValid(testCase *testing.T) {

	testingString := "{}()[]"
	result := queue.IsValid(testingString)
	expectedResult := true

	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	testingString = "((({[]})))"
	result = queue.IsValid(testingString)
	expectedResult = true

	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	testingString = "{}((("
	result = queue.IsValid(testingString)
	expectedResult = false

	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
}

func TestIsValid(testCase *testing.T) {
	testCase.Run("action=test-is-opener", testIsOpener)
	testCase.Run("action=test-is-valid", testIsValid)
}
