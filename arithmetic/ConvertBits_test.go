package arithmetic_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/arithmetic"
)

func testConvertAndReverse(testCase *testing.T) {
	var testNum uint32 = 12
	var expectedResult uint32 = 805306368
	var result uint32 = arithmetic.ConvertAndReverse(testNum)
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	testNum = 43261596
	expectedResult = 964176192
	result = arithmetic.ConvertAndReverse(testNum)
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
}

func TestConvertBits(testCase *testing.T) {
	testCase.Run("action=test-convert-and-reverse", testConvertAndReverse)
}
