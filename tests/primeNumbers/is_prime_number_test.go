package primeNumbers_test

import (
	"testing"

	"github.com/mrpiggy97/golang-leetcode/primeNumbers"
)

func TestIsPrimeNumber(testCase *testing.T) {
	var number int = 1231235
	var secondNumber int = 53
	var isPrimeNumber = primeNumbers.IsPrimeNumber(number)
	var shouldBePrimeNumber bool = primeNumbers.IsPrimeNumber((secondNumber))
	if isPrimeNumber != false {
		testCase.Error("expected isPrimeNumber to be false")
	}

	if !shouldBePrimeNumber {
		testCase.Error("shouldBePrimeNumber is expected to be true")
	}
}
