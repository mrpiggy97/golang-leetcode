package arithmetic

import "fmt"

func DivisorFind(number int) int {
	var originalNumber int = number
	var currentDivisor int = 2
	var totalNumberOfDivisors int = 1
	var numberOfDivisorInstances int = 0

	if number == 1 {
		return 1
	}
	for number > 1 {
		for number%currentDivisor == 0 {
			numberOfDivisorInstances++
			number = number / currentDivisor
		}
		if currentDivisor == 2 {
			currentDivisor++
		} else {
			currentDivisor += 2
		}
		totalNumberOfDivisors = totalNumberOfDivisors * (numberOfDivisorInstances + 1)
		numberOfDivisorInstances = 0
	}
	fmt.Printf("%v divisors for %v\n", totalNumberOfDivisors, originalNumber)
	return totalNumberOfDivisors
}
