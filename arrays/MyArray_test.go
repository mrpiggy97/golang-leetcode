package arrays_test

import (
	"fmt"
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

func testInsertAt(testCase *testing.T) {
	var testingArray *arrays.Array[int] = arrays.NewArray[int]()
	var first *int = new(int)
	var second *int = new(int)
	var fourth *int = new(int)
	var fifth *int = new(int)
	*first = 1
	*second = 2
	*fourth = 4
	*fifth = 5

	var third *int = new(int)
	*third = 3

	testingArray.Append(first)
	testingArray.Append(second)
	testingArray.Append(fourth)
	testingArray.Append(fifth)
	testingArray.InsertAt(2, third)
	var random *int = new(int)
	*random = 100
	testingArray.InsertAt(4, random)

	var storedValue int = *testingArray.Get(4)

	if storedValue != *random {
		testCase.Errorf("expected %d to be %d", storedValue, *third)
	}
	var seventh *int = new(int)
	*seventh = 7
	testingArray.InsertAt(0, seventh)
	if *testingArray.Get(0) != 7 {
		testCase.Errorf("expected %d, got %d instead", *seventh, *testingArray.Get(0))
	}
}

func testGetSlice(testCase *testing.T) {
	var testingArray *arrays.Array[int] = arrays.NewArray[int]()
	var first *int = new(int)
	var second *int = new(int)
	var fourth *int = new(int)
	var fifth *int = new(int)
	var third *int = new(int)
	*first = 1
	*second = 2
	*fourth = 4
	*fifth = 5
	*third = 100
	testingArray.Append(first)
	testingArray.Append(second)
	testingArray.Append(fourth)
	testingArray.Append(fifth)
	testingArray.InsertAt(2, third)

	var expectedResult []int = []int{*first, *second, *third, *fourth, *fifth}
	var result []int = testingArray.GetSlice()

	for index, value := range result {
		var currentComparison int = expectedResult[index]
		if currentComparison != value {
			testCase.Errorf("%d not equal to %d", value, currentComparison)
		}
	}
	fmt.Println(result)
}

func TestMyArray(testCase *testing.T) {
	testCase.Run("action=basic-functionality", testBasicFunctionality)
	testCase.Run("action=test-insert-at-start", testInsertAtStart)
	testCase.Run("action=test-insert-at", testInsertAt)
	testCase.Run("action=test-get-slice", testGetSlice)
}
