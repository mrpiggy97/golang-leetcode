package arrays_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/arrays"
)

func testFindSwap(testCase *testing.T) {
	var slice1 []int = []int{1, 2, 1, 2}
	var slice2 []int = []int{2, 2, 1, 1}
	var expectedResult bool = true
	var result bool = arrays.FindSwap(slice1, slice2, 0, 2)
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
	expectedResult = true
	result = arrays.FindSwap(slice2, slice1, 0, 1)
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	slice1 = []int{1, 1, 2, 2}
	slice2 = []int{1, 2, 2, 1}
	result = arrays.FindSwap(slice1, slice2, 1, 2)
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	slice1 = []int{1, 2, 2}
	slice2 = []int{2, 2, 1}
	expectedResult = true
	result = arrays.FindSwap(slice1, slice2, 0, 2)
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}

	slice1 = []int{1, 2, 2}
	slice2 = []int{2, 1, 1}
	expectedResult = false
	result = arrays.FindSwap(slice1, slice2, 0, 2)
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
	result = arrays.FindSwap(slice2, slice1, 0, 1)
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
}

func testArraysAreSimilar(testCase *testing.T) {

	slice1 := []int{1, 2, 2}
	slice2 := []int{2, 1, 1}
	expectedResult := false
	result := arrays.ArraysAreSimilar(slice1, slice2)
	if result != expectedResult {
		testCase.Errorf("expected %v to be %v", result, expectedResult)
	}
}

func TestArraysAreSimilar(testCase *testing.T) {
	testCase.Run("action=test-find-swap", testFindSwap)
	testCase.Run("action=test-arrays-are-similar", testArraysAreSimilar)
}
