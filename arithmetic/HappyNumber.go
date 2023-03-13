package arithmetic

import (
	"fmt"
	"math"
	"strconv"
)

// For every byte member of str convert that member to int
// and add that member**2 to result, and at the end return
// result as string
func CreateNewNumber(str string) string {
	var result int = 0
	for index := range str {
		currentNumber, _ := strconv.Atoi(string(str[index]))
		var power int = int(math.Pow(float64(currentNumber), 2))
		result += power
	}
	return fmt.Sprintf("%d", result)
}

func HappyNumber(number int) bool {
	var checker map[int]bool = make(map[int]bool)
	var currentStr string = fmt.Sprintf("%d", number)
	var index int = 0
	var current int = number
	var stop bool = false
	for !stop {
		if index == 1 {
			checker[current] = true
		} else {
			currentStr = CreateNewNumber(currentStr)
			current, _ = strconv.Atoi(currentStr)
			if current == 1 {
				break
			}
			if current > 1 {
				_, exists := checker[current]
				if exists {
					return false
				}
				checker[current] = true
			}
		}
	}
	return true
}
