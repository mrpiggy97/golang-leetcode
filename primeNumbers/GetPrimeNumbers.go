package primeNumbers

import "fmt"

func IsPrimeNumber(number int) bool {

	if number == 2 || number == 3 || number == 5 || number == 7 {
		return true
	}

	if number%2 == 0 {
		return false
	}
	if number%3 == 0 {
		return false
	}

	if number%5 == 0 {
		return false
	}

	if number%7 == 0 {
		return false
	}

	if number%9 == 0 {
		return false
	}
	return true
}

func GetPrimeNumbers(number int) string {
	var results []map[int]int = []map[int]int{}
	var divisor int = 2
	for number != 1 {
		var isPrime bool = IsPrimeNumber(divisor)
		var divisorCounter int = 0
		if isPrime && number%divisor == 0 {
			for number%divisor == 0 {
				var visual string = fmt.Sprintf("%v/%v", number, divisor)
				fmt.Println(visual)
				number = number / divisor
				divisorCounter = divisorCounter + 1
			}
			var counter map[int]int = make(map[int]int)
			counter[divisor] = divisorCounter
			results = append(results, counter)
		}
		divisor = divisor + 1
	}
	var finalString string = ""
	for _, counterMap := range results {
		for key, value := range counterMap {
			if value > 1 {
				finalString = finalString + fmt.Sprintf("(%v**%v)", key, value)
			} else {
				finalString = finalString + fmt.Sprintf("(%v)", key)
			}
		}
	}
	return finalString
}
