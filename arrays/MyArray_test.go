package arrays_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/arrays"
)

func testBasicFunctionality(testCase *testing.T) {
	var firstArray *arrays.Array[int] = arrays.NewArray[int]()
	var currentAge *int = new(int)
	*currentAge = 25
	var previousAge *int = new(int)
	*previousAge = 10
	firstArray.Append(currentAge)
	firstArray.Append(previousAge)
	var storedValue *int = firstArray.Get(0)
	var expectedResult int = *currentAge
	if expectedResult != *storedValue {
		testCase.Errorf("expected %d, got %d instead", expectedResult, *storedValue)
	}
	storedValue = firstArray.Get(1)
	expectedResult = 10
	if expectedResult != *storedValue {
		testCase.Errorf("expected %d, got %d instead", expectedResult, *storedValue)
	}
}

func testInsertAtStart(testCase *testing.T) {
	var secondArray *arrays.Array[int] = arrays.NewArray[int]()
	var currentYear *int = new(int)
	var previousYear *int = new(int)
	*currentYear = 2022
	*previousYear = 2021
	secondArray.Append(currentYear)
	secondArray.Append(previousYear)
	var milleniumYear *int = new(int)
	*milleniumYear = 2000
	secondArray.InsertAtStart(milleniumYear)

	var expectedResult int = 2000
	var storedValue int = *secondArray.Get(0)

	if expectedResult != storedValue {
		testCase.Errorf("expected %d, got %d instead", expectedResult, storedValue)
	}
	var previousCentury *int = new(int)
	*previousCentury = 1900
	secondArray.InsertAtStart(previousCentury)

	expectedResult = 1900
	storedValue = *secondArray.Get(0)
	if expectedResult != storedValue {
		testCase.Errorf("expected %d, got %d instead", expectedResult, storedValue)
	}
	var errorValue *int = secondArray.Get(1000)
	if errorValue != nil {
		testCase.Errorf("expected %d, got %d instead", expectedResult, storedValue)
	}
}

func TestMyArray(testCase *testing.T) {
	testCase.Run("action=basic-functionality", testBasicFunctionality)
	testCase.Run("action=test-insert-at-start", testInsertAtStart)
}
