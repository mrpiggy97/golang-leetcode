package arithmetic_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/arithmetic"
)

func testCreateNewNumber(testCase *testing.T) {
	var testNum string = "81"
	var expectedResult string = "65"
	var result string = arithmetic.CreateNewNumber(testNum)
	if result != expectedResult {
		testCase.Errorf("expected %s to be %s", result, expectedResult)
	}
}

func testHappyNumber(testCase *testing.T) {
	var testNum int = 2
	var expectedResult bool = false
	var result bool = arithmetic.HappyNumber(testNum)
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	testNum = 19
	expectedResult = true
	result = arithmetic.HappyNumber(testNum)
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	testNum = 7
	expectedResult = true
	result = arithmetic.HappyNumber(testNum)
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
}

func TestHappyNumber(testCase *testing.T) {
	testCase.Run("action=test-create-new-number", testCreateNewNumber)
	testCase.Run("action=test-happy-number", testHappyNumber)
}
