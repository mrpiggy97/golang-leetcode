package primeNumbers_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/primeNumbers"
)

func TestGetPrimeNumbers(testCase *testing.T) {
	var number int = 7775460
	var finalString string = primeNumbers.GetPrimeNumbers(number)
	if finalString != "(2**2)(3**3)(5)(7)(11**2)(17)" {
		testCase.Error("error with ", finalString)
	}
}
